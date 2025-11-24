<script setup lang="ts">
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import MainLayout from './layouts/MainLayout.vue'

const route = useRoute()
const useLayout = computed(() => route.meta.layout !== false)
</script>

<template>
  <MainLayout v-if="useLayout">
    <router-view />
  </MainLayout>
  <router-view v-else />
</template>
