package micro

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"
)

type (
	HTTPHandler func(http.ResponseWriter, *http.Request)

	Handler[TInput any, TOutput any] func(ctx Context[TInput]) (*TOutput, error)
)

func MakeHandler[TInput any, TOutput any](handler Handler[TInput, TOutput]) HTTPHandler {
	return func(writer http.ResponseWriter, request *http.Request) {
		// TODO: Enable this in production
		// Check if request body is application/json
		//if request.Header.Get("Content-Type") != "application/json" {
		//	writer.WriteHeader(http.StatusBadRequest)
		//	return
		//}

		input, err := decodeRequestBody[TInput](request)
		if err != nil {
			slog.Error("error decoding request body", "error", err)

			sendError(writer, http.StatusBadRequest, "invalid request body")
			return
		}

		err = ValidateStruct[TInput](&input)
		if err != nil {
			var validationError *ValidationError
			if errors.As(err, &validationError) {
				slog.Error("validation error", "error", err)

				sendError(writer, http.StatusUnprocessableEntity, "validation failed")
				return
			}

			slog.Error("error validating input", "error", err)

			sendError(writer, http.StatusInternalServerError, "internal server error")
			return
		}

		slog.Info("handling request", "requestId", request.Header.Get("X-Request-ID"))

		ctx := &handlerContext[TInput]{
			header: request.Header,
			input:  input,
		}

		output, err := handler(ctx)
		if err != nil {
			var serviceError *baseError
			if errors.As(err, &serviceError) {
				sendError(writer, serviceError.StatusCode, serviceError.Message)
				return
			}

			slog.Error("error handling request", "error", err)

			sendError(writer, http.StatusInternalServerError, "internal server error")
			return
		}

		sendJSON(writer, http.StatusOK, output)
	}
}

func decodeRequestBody[TInput any](request *http.Request) (TInput, error) {
	var input TInput
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&input)
	if err != nil && !errors.Is(err, io.EOF) {
		return input, err
	}

	return input, nil
}

func sendJSON[T any](writer http.ResponseWriter, statusCode int, data T) {
	payload, err := json.Marshal(data)
	if err != nil {
		slog.Error("error encoding data", "error", err)

		sendError(writer, http.StatusInternalServerError, "internal server error")
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	_, err = writer.Write(payload)
	if err != nil {
		slog.Error("error writing response", "error", err)

		sendError(writer, http.StatusInternalServerError, "internal server error")
		return
	}
}

func sendError(writer http.ResponseWriter, statusCode int, message string) {
	sendJSON(writer, statusCode, map[string]string{
		"error": message,
	})
}
