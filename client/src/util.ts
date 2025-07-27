import { parse } from "iso8601-duration";

export function parseTimeString(timeString: string): string {
  const { hours, minutes, seconds } = parse(timeString);

  const flooredSeconds = typeof seconds === "number" ? Math.floor(seconds) : 0;

  return `${hours}:${minutes}:${flooredSeconds}`;
}
