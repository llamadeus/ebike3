import { CombinedError } from "urql";


/**
 * Extracts the error message from a combined error.
 *
 * @param err The error to extract the message from.
 * @returns The error message.
 */
export function errorMessage(err: unknown): string {
  if (err instanceof CombinedError) {
    return err.graphQLErrors[0].message ?? "unknown error";
  }

  if (err instanceof Error) {
    return err.message;
  }

  return String(err);
}
