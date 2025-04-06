interface HasTypename {
  __typename?: string;
}

export type NonNullish<T> = Exclude<T, null | undefined>;
export type WithRequired<T, K extends keyof T> = T & { [P in K]-?: T[P] };
export type WithTypename<T extends HasTypename> = WithRequired<T, "__typename">
