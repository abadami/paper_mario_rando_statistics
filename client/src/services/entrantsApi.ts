import type { Entrant } from "../types";
import { API_CONFIG } from "../config";

export async function fetchEntrants(): Promise<Entrant[]> {
  try {
    const response = await fetch(
      `${API_CONFIG.BASE_URL}${API_CONFIG.ENDPOINTS.ENTRANTS}`,
    );

    if (!response.ok) {
      const errorText = await response.text();
      throw new Error(
        `Failed to fetch entrants: ${response.status} ${response.statusText}. ${errorText}`,
      );
    }

    const data = await response.json();
    return data;
  } catch (error) {
    if (error instanceof TypeError && error.message.includes("fetch")) {
      throw new Error(
        "Network error: Unable to connect to the entrants API. Please ensure the server is running.",
      );
    }
    throw error;
  }
}
