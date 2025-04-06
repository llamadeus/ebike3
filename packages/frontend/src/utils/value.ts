/**
 * Checks if a value is null or undefined.
 *
 * @param value The value to check.
 * @returns True if the value is null or undefined, false otherwise.
 */
export function isNullish(value: unknown): value is null | undefined {
  return typeof value == "undefined" || value === null;
}

/**
 * Checks if a value is not null or undefined.
 *
 * @param value The value to check.
 * @returns True if the value is not null or undefined, false otherwise.
 */
export function isNotNullish(value: unknown): value is NonNullable<typeof value> {
  return ! isNullish(value);
}

/**
 * Throws an error if the value is null or undefined.
 *
 * @param value The value to check.
 */
export function requireNullish(value: unknown): asserts value is null | undefined {
  if (isNullish(value)) {
    throw new Error("Should never happen");
  }
}

/**
 * Throws an error if the value is not null or undefined.
 *
 * @param value The value to check.
 */
export function requireNotNullish(value: unknown): asserts value is NonNullable<typeof value> {
  if (isNotNullish(value)) {
    throw new Error("Should never happen");
  }
}
