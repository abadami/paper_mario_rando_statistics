import { fetchStatistics } from "../data-controllers/statistic.controller";
import type { StatisticsResponse } from "../types";
import { updateAverageCard, updateDeviationCard } from "../ui-managers/card.ui";
import { updateRawDataTable } from "../ui-managers/raw-data-table.ui";
import { compareDesc } from "date-fns";

export async function updateStatistics(
  goal: string = "Blitz / 4 Chapters LCL Beat Bowser",
  entrant_id?: number,
  raceType?: string
) {
  const filters: { filter: string; value: string | number }[] = [];

  if (entrant_id) {
    filters.push({ filter: "ContainsEntrant", value: entrant_id });
  }

  if (goal) {
    filters.push({ filter: "Goal", value: goal });
  }

  if (raceType) {
    filters.push({ filter: "RaceType", value: raceType });
  }

  try {
    const statistics: StatisticsResponse = await fetchStatistics(filters);

    const sortRawData = statistics.rawData.sort((raceA, raceB) =>
      compareDesc(raceA.started_at, raceB.started_at)
    );

    updateAverageCard(statistics.average);

    updateDeviationCard(statistics.deviation);

    updateRawDataTable(sortRawData);
  } catch (error) {
    console.log(error);
    return;
  }
}
