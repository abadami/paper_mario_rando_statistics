<template>
  <div
    v-if="show"
    class="statistic-card bg-pm-dark border border-white rounded p-6"
  >
    <h2 class="statistic-title text-white text-xl font-semibold mb-2 m-0">
      {{ title }}
    </h2>
    <span class="statistic-value text-white text-2xl font-bold">
      {{ value }}
    </span>
  </div>
</template>

<script setup lang="ts">
interface Props {
  title: string
  value: string
  show: boolean
}

defineProps<Props>()
</script>

<style scoped>
.pm-dark {
  background-color: rgb(30, 49, 64);
}

.statistic-card {
  font-size: 18px;
}

.statistic-title {
  margin: 0;
  padding: 0;
}
</style>
