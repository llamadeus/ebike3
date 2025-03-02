import { GraphQLError } from "graphql/error";
import { z } from "zod";
import { Json } from "~/infrastructure/types/json";


type Service =
  | "auth"
  | "stations"
  | "vehicles";

type RequestHeaders = {
  "X-Request-ID": string,
  [key: string]: string,
};

interface GetOptions<TOutput extends z.ZodTypeAny> {
  endpoint: `GET /${string}`;
  headers: RequestHeaders;
  output: TOutput;
}

interface PostOptions<TOutput extends z.ZodTypeAny> {
  endpoint: `POST /${string}` | `PUT /${string}` | `PATCH /${string}` | `DELETE /${string}`;
  headers: RequestHeaders;
  input?: Json;
  output: TOutput;
}

type Options<TOutput extends z.ZodTypeAny> =
  | GetOptions<TOutput>
  | PostOptions<TOutput>;

const SERVICE_MAP: Record<Service, string> = {
  auth: "http://auth-service:5001",
  stations: "http://stations-service:5001",
  vehicles: "http://vehicles-service:5001",
};

export async function invokeService<TOutput extends z.ZodTypeAny>(
  service: Service,
  options: Options<TOutput>,
): Promise<z.infer<TOutput>> {
  const [method, endpoint] = options.endpoint.split(" ");
  if (typeof method == "undefined" || typeof endpoint == "undefined") {
    throw new Error("Invalid endpoint");
  }

  const url = new URL(endpoint, SERVICE_MAP[service]);
  const response = await fetch(url, {
    method: method,
    headers: {
      ...options.headers,
      "Content-Type": "application/json",
    },
    body: method !== "GET" && "input" in options
      ? JSON.stringify(options.input)
      : undefined,
  });

  let data: Json;
  try {
    data = await response.clone().json();
  }
  catch {
    const text = await response.text();
    throw new Error(text);
  }

  if (! response.ok) {
    const error = typeof data == "object" && data !== null && "error" in data && typeof data.error == "string"
      ? data.error
      : null;

    throw new GraphQLError(error ?? "Unknown error");
  }

  const output = options.output.safeParse(data);
  if (! output.success) {
    throw new GraphQLError(`Invalid response format: ${output.error.message}`);
  }

  return output.data;
}
