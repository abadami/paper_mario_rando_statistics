<template>
    <div class="bg-pm-dark border border-white rounded p-6">
        <h3 class="text-white text-xl font-semibold mb-4">Racer's Raw Data</h3>

        <!-- Empty State -->
        <div v-if="!data || data.length === 0" class="text-center py-8">
            <p class="text-white text-lg">No race data available.</p>
            <p class="text-gray-300 text-sm mt-2">
                Select a racer to view their race history.
            </p>
        </div>

        <!-- Data Table -->
        <DataTable
            v-else
            :value="data"
            :rows="10"
            :paginator="data.length > 10"
            :rowsPerPageOptions="[10, 20, 50]"
            class="p-datatable-sm"
            :pt="{
                root: 'bg-pm-dark',
                header: 'bg-pm-dark border-b border-white',
                tbody: 'bg-pm-dark',
                paginator: 'bg-pm-dark border-t border-white',
            }"
        >
            <Column field="started_at" header="Date" class="text-white">
                <template #body="{ data }">
                    <span class="text-white">
                        {{ formatDate(data.started_at) }}
                    </span>
                </template>
            </Column>

            <Column field="name" header="Race Name" class="text-white">
                <template #body="{ data }">
                    <a
                        :href="`https://racetime.gg/${data.name}`"
                        target="_blank"
                        class="text-blue-300 hover:text-blue-100 underline"
                    >
                        {{ data.name }}
                    </a>
                </template>
            </Column>

            <Column field="place_ordinal" header="Place" class="text-white">
                <template #body="{ data }">
                    <span class="text-white">
                        {{ data.status === "dnf" ? "--" : data.place_ordinal }}
                    </span>
                </template>
            </Column>

            <Column field="finish_time" header="Finish Time" class="text-white">
                <template #body="{ data }">
                    <span class="text-white">
                        {{
                            data.status === "dnf"
                                ? "DNF"
                                : parseTimeString(data.finish_time)
                        }}
                    </span>
                </template>
            </Column>
        </DataTable>
    </div>
</template>

<script setup lang="ts">
import DataTable from "primevue/datatable";
import Column from "primevue/column";
import { parseTimeString } from "../utils";
import type { RaceEntrantAndRaceRecord } from "../types";

interface Props {
    data: RaceEntrantAndRaceRecord[];
}

defineProps<Props>();

function formatDate(dateString: string): string {
    return new Date(dateString).toLocaleDateString();
}
</script>

<style scoped>
.pm-dark {
    background-color: rgb(30, 49, 64);
}

:deep(.p-datatable .p-datatable-header) {
    background-color: rgb(30, 49, 64);
    color: white;
    border-bottom: 1px solid white;
}

:deep(.p-datatable .p-datatable-thead > tr > th) {
    background-color: rgb(30, 49, 64);
    color: white;
    border: 1px solid white;
    font-size: 18px;
}

:deep(.p-datatable .p-datatable-tbody > tr > td) {
    background-color: rgb(30, 49, 64);
    color: white;
    border: 1px solid #555;
    font-size: 18px;
}

:deep(.p-datatable .p-datatable-tbody > tr:nth-child(even)) {
    background-color: rgb(25, 44, 59);
}

:deep(.p-datatable .p-datatable-tbody > tr:hover) {
    background-color: rgb(35, 54, 69);
}

:deep(.p-paginator) {
    background-color: rgb(30, 49, 64);
    color: white;
    border-top: 1px solid white;
}

:deep(.p-paginator .p-paginator-pages .p-paginator-page) {
    background-color: transparent;
    color: white;
    border: 1px solid white;
    margin: 0 2px;
}

:deep(.p-paginator .p-paginator-pages .p-paginator-page:hover) {
    background-color: rgba(255, 255, 255, 0.1);
}

:deep(.p-paginator .p-paginator-pages .p-paginator-page.p-highlight) {
    background-color: #3283b4;
    color: white;
}

:deep(.p-dropdown-panel) {
    background-color: rgb(30, 49, 64);
    border: 1px solid white;
}

:deep(.p-dropdown-panel .p-dropdown-items .p-dropdown-item) {
    color: white;
}

:deep(.p-dropdown-panel .p-dropdown-items .p-dropdown-item:hover) {
    background-color: rgba(255, 255, 255, 0.1);
}
</style>
