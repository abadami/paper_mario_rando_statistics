import type { RaceEntrantAndRaceRecord, StatisticsResponse } from "./types";
import { format } from "date-fns";
import { parseTimeString } from "./util";

function updateAverageCard(value: string) {
  const element = document.querySelector<HTMLSpanElement>("#average-value")!;

  element.textContent = value;
}

function updateDeviationCard(value: string) {
  const element = document.querySelector<HTMLSpanElement>("#deviation-value")!;

  element.textContent = value;
}

function updateRawDataTable(records: RaceEntrantAndRaceRecord[]) {
  const element = document.querySelector<HTMLTableSectionElement>(
    "#statistics-raw-data"
  )!;

  for (const record of records) {
    const newRow = element.insertRow();

    newRow.insertCell().textContent = new Date(
      record.started_at
    ).toLocaleDateString();
    newRow.insertCell().textContent = `${record.name}`;
    newRow.insertCell().textContent =
      record.status === "dnf" ? "--" : record.place_ordinal;
    newRow.insertCell().textContent =
      record.status === "dnf" ? "DNF" : parseTimeString(record.finish_time);
  }
}

export async function updateStatistics(entrant_id?: number) {
  const filters: { filter: string; value: string | number }[] = [];

  if (entrant_id) {
    filters.push({ filter: "ContainsEntrant", value: entrant_id });
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
