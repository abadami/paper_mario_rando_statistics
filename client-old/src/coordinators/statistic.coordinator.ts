import { fetchStatistics } from "../data-controllers/statistic.controller";
import type { StatisticsResponse } from "../types";
import { updateCard } from "../ui-managers/card.ui";
import { updateRawDataTable } from "../ui-managers/raw-data-table.ui";
import { compareDesc } from "date-fns";
import { disableLoading, enableLoading } from "../ui-managers/ui-manager.ui";

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
    enableLoading();

    const statistics: StatisticsResponse = await fetchStatistics(filters);

    const data =
      typeof entrant_id === "undefined" || entrant_id === -1
        ? statistics.fullRawData
        : statistics.rawData;

    const sortRawData = data.sort((raceA, raceB) =>
      compareDesc(raceA.started_at, raceB.started_at)
    );

    updateCard("#average-value", statistics.average);
    updateCard("#deviation-value", statistics.deviation);
    updateCard("#average-win-value", statistics.averageWin);
    updateCard("#best-win-value", statistics.bestWin);
    updateCard("#worse-loss-value", statistics.worstLoss);

    updateRawDataTable(sortRawData);

    disableLoading();
  } catch (error) {
    disableLoading();
    console.log(error);
    return;
  }
}
