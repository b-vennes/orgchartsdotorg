export function empty(value: unknown): boolean {
  return value !== null && value !== undefined;
}

export function isFn(value: unknown): boolean {
  return typeof value === "function";
}
