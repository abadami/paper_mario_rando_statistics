import { fetchCategoryGoals } from "../data-controllers/category.controller";
import { updateStatistics } from "../coordinators/statistic.coordinator";
import type { Entrant } from "../types";
import { fetchEntrants } from "../data-controllers/entrant.controller";

let entrant_id = -1;
let category = "Blitz / 4 Chapters LCL Beat Bowser";

export async function setupUserFilter(element: HTMLSelectElement) {
  try {
    const entrants: Entrant[] = await fetchEntrants();

    element.innerHTML += `<option value=-1>All</option>`;

    for (const entrant of entrants) {
      element.innerHTML += `<option value=${entrant.id}>${entrant.name}</option>`;
    }

    element.addEventListener("change", (e: Event) => {
      const elementValue = (e.target as HTMLSelectElement).value;

      const value = parseInt(elementValue);

      entrant_id = value;

      updateStatistics(category, value);
    });
  } catch (error) {
    console.log(error);
    return;
  }
}

export async function setupCategoryFilter(element: HTMLSelectElement) {
  try {
    const goals = await fetchCategoryGoals();

    for (const goal of goals) {
      element.innerHTML += `<option value="${goal}">${goal}</option>`;
    }

    element.value = category;

    element.addEventListener("change", (e: Event) => {
      const goal = (e.target as HTMLSelectElement).value;

      category = goal;

      updateStatistics(goal, entrant_id);
    });
  } catch (error) {
    console.log("How did we get here");
    console.log(error);
    return;
  }
}
