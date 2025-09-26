export function empty<A>(value: A | undefined | null): value is undefined | null {
  return value === undefined || value === null;
}

export function nonempty<A>(value: A | undefined | null): value is A {
  return value !== null && value !== undefined;
}

export function isFn(value: unknown): boolean {
  return typeof value === "function";
}

export function hasLength<A>(value: unknown): value is Array<unknown> {
  const valueAsArray = value as Array<unknown>;

  return typeof valueAsArray === "object" &&
    nonempty(valueAsArray) &&
    nonempty(valueAsArray.length);
}
