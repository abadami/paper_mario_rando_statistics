import type { StatisticsResponse } from "./types";
import {
  updateAverageCard,
  updateDeviationCard,
} from "./ui-managers/card-manager.ui";
import { updateRawDataTable } from "./ui-managers/raw-data-table-manager.ui";

export async function updateStatistics(
  goal: string = "Blitz / 4 Chapters LCL Beat Bowser",
  entrant_id?: number
) {
  const filters: { filter: string; value: string | number }[] = [];

  if (entrant_id) {
    filters.push({ filter: "ContainsEntrant", value: entrant_id });
  }

  if (goal) {
    filters.push({ filter: "Goal", value: goal });
  }

  try {
    const response = await fetch(
      `http://localhost:3000/api/get_statistics_for_entrant${
        filters.length > 0 ? "?" : ""
      }${filters.map((filter) => `${filter.filter}=${filter.value}`).join("&")}`
    );
    const statistics: StatisticsResponse = await response.json();

    updateAverageCard(statistics.average);

    updateDeviationCard(statistics.deviation);

    updateRawDataTable(statistics.rawData);
  } catch (error) {
    console.log(error);
    return;
  }
}
