import { useQuery } from "@tanstack/vue-query";
import { fetchCategoryGoals } from "../services/categoriesApi";
import { QUERY_CONFIG } from "../config";

export function useCategories() {
  return useQuery({
    queryKey: ["categories"],
    queryFn: fetchCategoryGoals,
    staleTime: QUERY_CONFIG.STALE_TIME.CATEGORIES,
    gcTime: QUERY_CONFIG.GC_TIME.CATEGORIES,
  });
}
