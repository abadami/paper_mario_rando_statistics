<script setup lang="ts">
defineProps<{
  title?: string
  value: number | string
}>()
</script>

<template>
  <div class="statistic-block">
    <h4 class="statistic-header">{{ title }}</h4>
    <p class="statistic-value">{{ value }}</p>
  </div>
</template>

<style></style>
