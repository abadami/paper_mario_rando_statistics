import './style.css'
import { setupUserFilter } from './filters.ts'


await setupUserFilter(document.querySelector<HTMLSelectElement>("#user-selector")!)