import { updateStatistics } from "./statistic-controller";
import type { Entrant } from "./types";

export async function setupUserFilter(element: HTMLSelectElement) {
  try {
    const response = await fetch("http://localhost:3000/api/get_race_entrants");
    const entrants: Entrant[] = await response.json();

    element.innerHTML += `<option value="all">All</option>`;

    for (const entrant of entrants) {
      element.innerHTML += `<option value=${entrant.id}>${entrant.name}</option>`;
    }

    element.addEventListener("change", (e: Event) => {
      const elementValue = (e.target as HTMLSelectElement).value;

      console.log(elementValue);

      const value = parseInt((e.target as HTMLSelectElement).value);

      console.log("Selected Value: ", value);

      updateStatistics(value);
    });
  } catch (error) {
    console.log(error);
    return;
  }
}

export function setupCategoryFilter(element: HTMLSelectElement) {}
