<template>
    <main class="min-h-screen bg-pm-blue p-8">
        <div class="max-w-7xl mx-auto">
            <h1 class="text-4xl font-bold text-white text-center mb-8">
                Paper Mario 64 Randomizer Race Statistics
            </h1>

            <!-- Filters Section -->
            <div class="bg-pm-dark p-6 rounded border border-white mb-8">
                <h2 class="text-xl text-white mb-4">Statistics Filters</h2>
                <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
                    <!-- Racer Filter -->
                    <div class="flex flex-col">
                        <label for="racer-select" class="text-white mb-2"
                            >Racer</label
                        >
                        <Dropdown
                            id="racer-select"
                            v-model="selectedEntrant"
                            :options="entrantOptions"
                            option-label="name"
                            option-value="id"
                            placeholder="Select a racer"
                            class="w-full"
                            :loading="entrantsQuery.isPending.value"
                        />
                    </div>

                    <!-- Race Goal Filter -->
                    <div class="flex flex-col">
                        <label for="goal-select" class="text-white mb-2"
                            >Race Goal</label
                        >
                        <Dropdown
                            id="goal-select"
                            v-model="selectedGoal"
                            :options="goalOptions"
                            placeholder="Select a goal"
                            class="w-full"
                            :loading="categoriesQuery.isPending.value"
                        />
                    </div>

                    <!-- Race Type Filter -->
                    <div class="flex flex-col">
                        <label for="race-type-select" class="text-white mb-2"
                            >Racer Number</label
                        >
                        <Dropdown
                            id="race-type-select"
                            v-model="selectedRaceType"
                            :options="raceTypeOptions"
                            option-label="label"
                            option-value="value"
                            placeholder="Select race type"
                            class="w-full"
                        />
                    </div>
                </div>
            </div>

            <!-- Loading Section -->
            <div
                v-if="statisticsQuery.isPending.value"
                class="bg-pm-dark p-6 rounded border border-white mb-8"
            >
                <div class="flex items-center justify-center">
                    <ProgressSpinner class="mr-3" size="small" />
                    <span class="text-white">Loading...</span>
                </div>
            </div>

            <!-- Error Section -->
            <div
                v-if="statisticsQuery.isError.value"
                class="bg-red-900 p-6 rounded border border-red-500 mb-8"
            >
                <p class="text-white">
                    Error loading statistics:
                    {{ statisticsQuery.error.value?.message }}
                </p>
            </div>

            <!-- Statistics Display -->
            <div
                v-if="processedStatistics && !statisticsQuery.isPending.value"
                class="space-y-8"
            >
                <!-- Statistics Cards -->
                <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
                    <!-- Average Card -->
                    <StatisticCard
                        title="Average"
                        :value="processedStatistics.average"
                        :show="true"
                    />

                    <!-- Standard Deviation Card -->
                    <StatisticCard
                        title="Standard Deviation"
                        :value="processedStatistics.deviation"
                        :show="true"
                    />

                    <!-- Average Win Card -->
                    <StatisticCard
                        title="Average Win"
                        :value="processedStatistics.averageWin"
                        :show="showEntrantSpecificStats"
                    />

                    <!-- Best Win Card -->
                    <StatisticCard
                        title="Best Win"
                        :value="processedStatistics.bestWin"
                        :show="showLeagueStats"
                    />

                    <!-- Worst Loss Card -->
                    <StatisticCard
                        title="Worst Loss"
                        :value="processedStatistics.worstLoss"
                        :show="showLeagueStats"
                    />
                </div>

                <!-- Raw Data Table -->
                <RawDataTable
                    v-if="showRawDataTable"
                    :data="processedStatistics.sortedRawData"
                />
            </div>
        </div>
    </main>
</template>

<script setup lang="ts">
import { ref, computed, watch } from "vue";
import Dropdown from "primevue/dropdown";
import ProgressSpinner from "primevue/progressspinner";
import { useEntrants } from "./composables/useEntrants";
import { useCategories } from "./composables/useCategories";
import { useStatisticsData } from "./composables/useStatistics";
import StatisticCard from "./components/StatisticCard.vue";
import RawDataTable from "./components/RawDataTable.vue";
import type { Filter } from "./types";
import { API_CONFIG, RACE_TYPE_OPTIONS } from "./config";

// Reactive state
const selectedEntrant = ref<number>(-1);
const selectedGoal = ref<string>(API_CONFIG.DEFAULT_GOAL);
const selectedRaceType = ref<string>("");

// Race type options
const raceTypeOptions = [
    { label: "All", value: "" },
    { label: "League (1v1)", value: "league" },
    { label: "Community (More than 2 racers)", value: "community" },
];

// API queries
const entrantsQuery = useEntrants();
const categoriesQuery = useCategories();

// Computed options for dropdowns
const entrantOptions = computed(() => {
    const options = [{ id: -1, name: "All", full_name: "All" }];
    if (entrantsQuery.data.value) {
        options.push(...entrantsQuery.data.value);
    }
    return options;
});

const goalOptions = computed(() => {
    return categoriesQuery.data.value || [];
});

// Computed filters for statistics query
const statisticsFilters = computed<Filter[]>(() => {
    const filters: Filter[] = [];

    if (selectedEntrant.value && selectedEntrant.value !== -1) {
        filters.push({
            filter: "ContainsEntrant",
            value: selectedEntrant.value,
        });
    }

    if (selectedGoal.value) {
        filters.push({ filter: "Goal", value: selectedGoal.value });
    }

    if (selectedRaceType.value) {
        filters.push({ filter: "RaceType", value: selectedRaceType.value });
    }

    return filters;
});

// Statistics query
const statisticsQuery = useStatisticsData(statisticsFilters);

// Computed statistics data
const processedStatistics = computed(() => {
    return statisticsQuery.data.value;
});

// Computed visibility flags
const showEntrantSpecificStats = computed(() => {
    return selectedEntrant.value !== -1;
});

const showLeagueStats = computed(() => {
    return selectedEntrant.value !== -1 && selectedRaceType.value === "league";
});

const showRawDataTable = computed(() => {
    return (
        selectedEntrant.value !== -1 &&
        processedStatistics.value?.sortedRawData.length > 0
    );
});

// Watch for initial data load and set default goal
watch(
    () => categoriesQuery.data.value,
    (goals) => {
        if (goals && goals.length > 0 && !selectedGoal.value) {
            selectedGoal.value =
                goals[0] || "Blitz / 4 Chapters LCL Beat Bowser";
        }
    },
    { immediate: true },
);
</script>

<style scoped>
.pm-blue {
    background-color: #3283b4;
}

.pm-dark {
    background-color: rgb(30, 49, 64);
}
</style>
