import "./style.css";
import {
  setupUserFilter,
  setupCategoryFilter,
  setupRaceTypeFilter,
} from "./ui-managers/filter.ui.ts";
import { updateStatistics } from "./coordinators/statistic.coordinator.ts";

setupUserFilter(document.querySelector<HTMLSelectElement>("#user-selector")!);

setupCategoryFilter(
  document.querySelector<HTMLSelectElement>("#category-selector")!
);

setupRaceTypeFilter(
  document.querySelector<HTMLSelectElement>("#race-type-selector")!
);

updateStatistics();
