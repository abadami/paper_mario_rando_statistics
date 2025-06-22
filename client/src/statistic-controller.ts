import type { RaceEntrantAndRaceRecord, StatisticsResponse } from "./types";

function updateAverageCard(value: string) {
  const element = document.querySelector<HTMLSpanElement>("#average-value")!;

  element.textContent = value;
}

export async function updateStatistics(entrant_id: number) {
  try {
    const response = await fetch(
      `http://localhost:3000/api/get_statistics_for_entrant?ContainsEntrant=${entrant_id}`
    );
    const statistics: StatisticsResponse = await response.json();

    updateAverageCard(statistics.average);
  } catch (error) {
    console.log(error);
    return;
  }
}
