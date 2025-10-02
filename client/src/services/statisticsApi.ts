import type { StatisticsResponse, Filter } from "../types";
import { API_CONFIG } from "../config";

export async function fetchStatistics(
  filters: Filter[],
): Promise<StatisticsResponse> {
  const queryParams =
    filters.length > 0
      ? "?" +
        filters
          .map(
            (filter) => `${filter.filter}=${encodeURIComponent(filter.value)}`,
          )
          .join("&")
      : "";

  try {
    const response = await fetch(
      `${API_CONFIG.BASE_URL}${API_CONFIG.ENDPOINTS.STATISTICS}${queryParams}`,
    );

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(
        `Failed to fetch statistics: ${response.status} ${response.statusText}. ${errorText}`,
      );
    }

    const data = await response.json();
    return data;
  } catch (error) {
    if (error instanceof TypeError && error.message.includes("fetch")) {
      throw new Error(
        "Network error: Unable to connect to the statistics API. Please ensure the server is running.",
      );
    }
    throw error;
  }
}
