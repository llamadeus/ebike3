package micro

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"slices"
	"strings"
	"time"
)

var methods = []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete}

type InvokeError struct {
	Status  int
	Message string
}

func (e InvokeError) Error() string {
	return fmt.Sprintf("failed to invoke endpoint: %s (status %d)", e.Message, e.Status)
}

func Run(mux *http.ServeMux, serverAddr string) (err error) {
	server := &http.Server{
		Addr:    serverAddr,
		Handler: mux,
	}

	// Channel to listen for interrupt signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		slog.Info("starting server",
			"addr", serverAddr,
		)

		if err = server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("server failed", "error", err)
			close(stop)
		}
	}()

	// Wait for interrupt signal
	<-stop

	// Create a handlerContext with a timeout for shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	return server.Shutdown(ctx)
}

func Invoke[TInput any, TOutput any](ctx context.Context, endpoint string, requestID *string, input TInput) (TOutput, error) {
	var httpClient http.Client
	var output TOutput

	segments := strings.Split(endpoint, " ")
	if len(segments) != 2 {
		return output, fmt.Errorf("invalid endpoint: %s", endpoint)
	}

	method := segments[0]
	url := segments[1]

	if slices.Contains(methods, method) == false {
		return output, fmt.Errorf("invalid method: %s", method)
	}

	req, err := http.NewRequest(method, fmt.Sprintf("http://%s", url), nil)
	if err != nil {
		return output, err
	}

	// Request with context
	req = req.WithContext(ctx)

	if requestID != nil {
		req.Header.Set("X-Request-ID", *requestID)
	}

	t := reflect.TypeOf(input)
	if t != nil {
		if t.Kind() != reflect.Struct {
			return output, errors.New("TInput must be a struct")
		}

		payload, err := json.Marshal(input)
		if err != nil {
			return output, err
		}

		req.Header.Set("Content-Type", "application/json")
		req.Body = io.NopCloser(bytes.NewReader(payload))
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return output, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		type errorResponse struct {
			Error string `json:"error"`
		}

		var er errorResponse
		err = json.NewDecoder(resp.Body).Decode(&er)
		if err != nil {
			return output, err
		}

		slog.Error(
			"failed to invoke endpoint",
			"endpoint", endpoint,
			"status", resp.Status,
			"body", er.Error,
		)
		return output, &InvokeError{resp.StatusCode, er.Error}
	}

	err = json.NewDecoder(resp.Body).Decode(&output)

	return output, err
}
