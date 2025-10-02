export const API_CONFIG = {
  BASE_URL: 'http://localhost:3000/api',
  RACETIME_URL: 'https://racetime.gg',
  DEFAULT_GOAL: 'Blitz / 4 Chapters LCL Beat Bowser',
  ENDPOINTS: {
    STATISTICS: '/get_statistics_for_entrant',
    ENTRANTS: '/get_race_entrants',
    CATEGORY_DATA: 'https://racetime.gg/pm64r/data',
  },
} as const

export const QUERY_CONFIG = {
  STALE_TIME: {
    STATISTICS: 5 * 60 * 1000, // 5 minutes
    ENTRANTS: 15 * 60 * 1000, // 15 minutes
    CATEGORIES: 30 * 60 * 1000, // 30 minutes
  },
  GC_TIME: {
    STATISTICS: 10 * 60 * 1000, // 10 minutes
    ENTRANTS: 30 * 60 * 1000, // 30 minutes
    CATEGORIES: 60 * 60 * 1000, // 1 hour
  },
} as const

export const RACE_TYPE_OPTIONS = [
  { label: 'All', value: '' },
  { label: 'League (1v1)', value: 'league' },
  { label: 'Community (More than 2 racers)', value: 'community' },
] as const
