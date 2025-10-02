import type { StatisticsResponse } from "../types";

interface Filter {
  filter: string;
  value: string | number;
}

export async function fetchStatistics(
  filters: Filter[]
): Promise<StatisticsResponse> {
  const response = await fetch(
    `http://localhost:3000/api/get_statistics_for_entrant${
      filters.length > 0 ? "?" : ""
    }${filters.map((filter) => `${filter.filter}=${filter.value}`).join("&")}`
  );
  const statistics: StatisticsResponse = await response.json();

  return statistics;
}
