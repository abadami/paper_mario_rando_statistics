export interface Entrant {
  id: number;
  name: string;
  full_name: string;
}

export interface RaceEntrantAndRaceRecord {
  id: number;
  race_id: number;
  entrant_id: number;
  finish_time: string;
  place: number;
  place_ordinal: string;
  status: string;
  name: string;
  category_name: string;
  category_short_name: string;
  url: string;
  goal_name: string;
  started_at: string;
}

export interface StatisticsResponse {
  average: string;
  deviation: string;
  raceNumber: number;
  dnfCount: number;
  rawData: RaceEntrantAndRaceRecord[];
}
