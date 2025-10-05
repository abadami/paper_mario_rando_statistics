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
  bestWin: string;
  worstLoss: string;
  averageWin: string;
  raceNumber: number;
  dnfCount: number;
  rawData: RaceEntrantAndRaceRecord[];
  fullRawData: RaceEntrantAndRaceRecord[];
}

export interface CategoryDetails {
  name: string;
  short_name: string;
  slug: string;
  url: string;
  data_url: string;
  image: string;
  info: string;
  streaming_required: boolean;
  goals: string[];
}

export interface Filter {
  filter: string;
  value: string | number;
}
