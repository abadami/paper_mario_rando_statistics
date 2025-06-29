import { parse } from "iso8601-duration";

export function parseTimeString(timeString: string): string {
  const { hours, minutes, seconds } = parse(timeString);

  return `${hours}:${minutes}:${seconds}`;
}
