import { useQuery } from "@tanstack/vue-query";
import { computed, type Ref } from "vue";
import { fetchStatistics } from "../services/statisticsApi";
import type { Filter, StatisticsResponse } from "../types";
import { QUERY_CONFIG } from "../config";

export function useStatistics(filters: Ref<Filter[]>) {
  return useQuery({
    queryKey: ["statistics", filters],
    queryFn: () => fetchStatistics(filters.value),
    enabled: computed(() => filters.value.length > 0),
    staleTime: QUERY_CONFIG.STALE_TIME.STATISTICS,
    gcTime: QUERY_CONFIG.GC_TIME.STATISTICS,
  });
}

export function useStatisticsData(filters: Ref<Filter[]>) {
  const { data, isLoading, error, refetch, isPending, isError } =
    useStatistics(filters);

  const processedData = computed(() => {
    if (!data.value) return null;

    const statistics = data.value as StatisticsResponse;

    // Determine which data to use based on filters
    const hasEntrantFilter = filters.value.some(
      (f) => f.filter === "ContainsEntrant" && f.value !== -1,
    );
    const rawData = hasEntrantFilter
      ? statistics.rawData
      : statistics.fullRawData;

    // Sort by date descending (using date-fns compareDesc)
    const sortedData = [...rawData].sort((a, b) => {
      const dateA = new Date(a.started_at);
      const dateB = new Date(b.started_at);
      return dateB.getTime() - dateA.getTime();
    });

    return {
      ...statistics,
      sortedRawData: sortedData,
      hasEntrantFilter,
    };
  });

  return {
    data: processedData,
    isLoading,
    isPending,
    isError,
    error,
    refetch,
  };
}
