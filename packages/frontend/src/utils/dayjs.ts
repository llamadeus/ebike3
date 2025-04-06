import dayjs from "dayjs";


interface Difference {
  days: number;
  hours: number;
  minutes: number;
  seconds: number;
}

/**
 * Computes the difference between two dates.
 * Returns the number of days, hours, minutes, and seconds.
 *
 * @param start The start date.
 * @param end The end date.
 * @returns The difference between the two dates.
 */
export function computeDifference(start: Date, end: Date): Difference {
  const diff = dayjs(end).diff(start, "second");
  const days = Math.floor(diff / (3600 * 24));
  const hours = Math.floor((diff % (3600 * 24)) / 3600);
  const minutes = Math.floor((diff % 3600) / 60);
  const seconds = diff % 60;

  return {
    days,
    hours,
    minutes,
    seconds,
  };
}

/**
 * Formats the difference between two dates as a string in a human-readable format.
 *
 * @param start The start date.
 * @param end The end date.
 * @returns The formatted difference.
 */
export function formatDifference(start: Date, end: Date): string {
  const { days, hours, minutes, seconds } = computeDifference(start, end);
  const parts: string[] = [];

  if (days > 0) {
    parts.push(`${days}d`);
  }
  if (hours > 0) {
    parts.push(`${hours}h`);
  }
  if (minutes > 0) {
    parts.push(`${minutes}m`);
  }
  if (seconds > 0) {
    parts.push(`${seconds}s`);
  }

  return parts.join(" ");
}
