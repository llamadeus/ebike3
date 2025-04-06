const PLURALS = {
  "item": "items",
} as const;

/**
 * Pluralizes a word based on the given count.
 *
 * @param word
 * @param count
 * @param withCount
 */
export function pluralize(word: keyof typeof PLURALS, count: number, withCount = true): string {
  const str = Math.abs(count) === 1
    ? word
    : PLURALS[word];

  return withCount ? `${count} ${str}` : str;
}
