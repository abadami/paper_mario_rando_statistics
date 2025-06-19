<script setup lang="ts">
defineProps<{
  title?: string
  value: number | string
}>()
</script>

<template>
  <Card>
    <template #title>{{ title }}</template>
    <template #content>{{ value }}</template>
  </Card>
</template>

<style></style>
