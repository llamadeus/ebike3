/**
 * Checks if a value is null or undefined.
 *
 * @param value The value to check
 */
export function isNullOrUndefined(value: unknown): value is null | undefined {
  return typeof value == "undefined" || value === null;
}
