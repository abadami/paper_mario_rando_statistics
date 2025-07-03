import './style.css'
import { setupUserFilter, setupCategoryFilter } from './ui-managers/filter.ui.ts'
import { updateStatistics } from './coordinators/statistic.coordinator.ts'

setupUserFilter(document.querySelector<HTMLSelectElement>("#user-selector")!)

setupCategoryFilter(document.querySelector<HTMLSelectElement>("#category-selector")!)

updateStatistics()