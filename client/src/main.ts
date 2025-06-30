import './style.css'
import { setupUserFilter } from './filters.ts'
import { updateStatistics } from './statistic-controller.ts'


setupUserFilter(document.querySelector<HTMLSelectElement>("#user-selector")!)

updateStatistics()