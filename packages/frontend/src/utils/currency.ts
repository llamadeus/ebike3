/**
 * Formats a number as a currency string.
 *
 * @param value The value to format.
 * @returns The formatted currency string.
 */
export function formatCurrency(value: number) {
  return `${(value / 100).toFixed(2)} €`;
}
