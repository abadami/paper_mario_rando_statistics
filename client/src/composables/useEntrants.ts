import { useQuery } from "@tanstack/vue-query";
import { fetchEntrants } from "../services/entrantsApi";
import { QUERY_CONFIG } from "../config";

export function useEntrants() {
  return useQuery({
    queryKey: ["entrants"],
    queryFn: fetchEntrants,
    staleTime: QUERY_CONFIG.STALE_TIME.ENTRANTS,
    gcTime: QUERY_CONFIG.GC_TIME.ENTRANTS,
  });
}
