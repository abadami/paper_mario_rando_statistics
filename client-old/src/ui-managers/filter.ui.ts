import { fetchCategoryGoals } from "../data-controllers/category.controller";
import { updateStatistics } from "../coordinators/statistic.coordinator";
import type { Entrant } from "../types";
import { fetchEntrants } from "../data-controllers/entrant.controller";
import { hideElement, showElement } from "./ui-manager.ui";

let entrant_id = -1;
let category = "Blitz / 4 Chapters LCL Beat Bowser";
let raceType = "";

//TODO: Setup method for showing correct stat blocks based on filters

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

      showStatisticsBasedOnFilter(value, raceType);

      updateStatistics(category, value, raceType);
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

      updateStatistics(goal, entrant_id, raceType);
    });
  } catch (error) {
    console.log("How did we get here");
    console.log(error);
    return;
  }
}

export function setupRaceTypeFilter(element: HTMLSelectElement) {
  element.value = "";

  element.addEventListener("change", (e: Event) => {
    const type = (e.target as HTMLSelectElement).value;

    raceType = type;

    showStatisticsBasedOnFilter(entrant_id, type);

    updateStatistics(category, entrant_id, type);
  });
}

export function showStatisticsBasedOnFilter(entrant: number, type: string) {
  hideAllStatistics(["average", "standard-deviation"]);

  if (entrant <= -1) {
    return;
  } else {
    showStatistics(["average-win", "statistic-raw-data-table"]);
  }

  if (type === "league") {
    showStatistics(["best-win", "worse-loss"]);
  }
}

export function hideStatistics(ids: string[]) {
  for (const id of ids) {
    const element = document.querySelector<HTMLElement>(`#${id}`)!;

    hideElement(element);
  }
}

export function showStatistics(ids: string[]) {
  for (const id of ids) {
    const element = document.querySelector<HTMLElement>(`#${id}`)!;

    showElement(element);
  }
}

export function hideAllStatistics(exclusionList: string[]) {
  const element = document.querySelector<HTMLDivElement>(
    `#statistic-information`
  )!;

  for (const child of element.children) {
    if (exclusionList.includes(child.id)) {
      continue;
    }

    hideElement(child as HTMLElement);
  }
}
