import { parse } from "iso8601-duration";

export function parseTimeString(timeString: string): string {
  const { hours, minutes, seconds } = parse(timeString);

  const flooredSeconds = typeof seconds === "number" ? Math.floor(seconds) : 0;

  const formattedHours = hours ? formatSingleDigitNumbers(hours) : 0;

  const formattedMinutes = minutes ? formatSingleDigitNumbers(minutes) : 0;

  const formattedSeconds = formatSingleDigitNumbers(flooredSeconds);

  return `${formattedHours}:${formattedMinutes}:${formattedSeconds}`;
}

export function formatSingleDigitNumbers(number: number): string {
  return number.toLocaleString("en-US", {
    minimumIntegerDigits: 2,
    useGrouping: false,
  });
}
