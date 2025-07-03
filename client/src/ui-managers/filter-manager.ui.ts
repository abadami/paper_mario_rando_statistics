import { fetchCategoryGoals } from "../data-controllers/category.controller";
import { updateStatistics } from "../statistic-controller";
import type { Entrant } from "../types";

let entrant_id = -1;
let category = "Blitz / 4 Chapters LCL Beat Bowser";

export async function setupUserFilter(element: HTMLSelectElement) {
  try {
    const response = await fetch("http://localhost:3000/api/get_race_entrants");
    const entrants: Entrant[] = await response.json();

    element.innerHTML += `<option value=-1>All</option>`;

    for (const entrant of entrants) {
      element.innerHTML += `<option value=${entrant.id}>${entrant.name}</option>`;
    }

    element.addEventListener("change", (e: Event) => {
      const elementValue = (e.target as HTMLSelectElement).value;

      const value = parseInt(elementValue);

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
      element.innerHTML += `<option value=${goal}>${goal}</option>`;
    }

    element.addEventListener("change", (e: Event) => {
      const goal = (e.target as HTMLSelectElement).value;

      category = goal;

      updateStatistics(goal, entrant_id);
    });
  } catch (error) {
    console.log(error);
    return;
  }
}
