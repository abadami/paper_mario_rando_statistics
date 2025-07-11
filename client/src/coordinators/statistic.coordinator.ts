import { fetchStatistics } from "../data-controllers/statistic.controller";
import type { StatisticsResponse } from "../types";
import { updateAverageCard, updateDeviationCard } from "../ui-managers/card.ui";
import { updateRawDataTable } from "../ui-managers/raw-data-table.ui";

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
    const statistics: StatisticsResponse = await fetchStatistics(filters);

    updateAverageCard(statistics.average);

    updateDeviationCard(statistics.deviation);

    updateRawDataTable(statistics.rawData);
  } catch (error) {
    console.log(error);
    return;
  }
}
