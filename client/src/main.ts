import './style.css'
import { setupUserFilter, setupCategoryFilter } from './ui-managers/filter-manager.ui.ts'
import { updateStatistics } from './statistic-controller.ts'


setupUserFilter(document.querySelector<HTMLSelectElement>("#user-selector")!)

setupCategoryFilter(document.querySelector<HTMLSelectElement>("#category-selector")!)

updateStatistics()