import { clsx, type ClassValue } from "clsx";
import { twMerge } from "tailwind-merge";


/**
 * Builds the class name from the given inputs.
 * Removes duplicate classes using `tailwind-merge`.
 *
 * @param inputs
 */
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}
