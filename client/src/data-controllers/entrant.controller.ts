import type { Entrant } from "../types";

export async function fetchEntrants(): Promise<Entrant[]> {
  const response = await fetch("http://localhost:3000/api/get_race_entrants");
  const entrants: Entrant[] = await response.json();

  return entrants;
}
