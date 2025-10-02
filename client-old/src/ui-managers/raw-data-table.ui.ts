import type { RaceEntrantAndRaceRecord } from "../types";
import { parseTimeString } from "../util";

function insertDataIntoTable(record: RaceEntrantAndRaceRecord) {
  const element = document.querySelector<HTMLTableSectionElement>(
    "#statistics-raw-data"
  )!;

  const newRow = element.insertRow();

  newRow.insertCell().textContent = new Date(
    record.started_at
  ).toLocaleDateString();
  newRow.insertCell().innerHTML = `<a href="https://racetime.gg/${record.name}" target="_blank">${record.name}</a>`;
  newRow.insertCell().textContent =
    record.status === "dnf" ? "--" : record.place_ordinal;
  newRow.insertCell().textContent =
    record.status === "dnf" ? "DNF" : parseTimeString(record.finish_time);
}

function clearTable() {
  const element = document.querySelector<HTMLTableSectionElement>(
    "#statistics-raw-data"
  )!;

  for (let i = element.rows.length - 1; i > -1; i--) {
    element.deleteRow(i);
  }
}

export function updateRawDataTable(records: RaceEntrantAndRaceRecord[]) {
  clearTable();

  for (const record of records) {
    insertDataIntoTable(record);
  }
}
