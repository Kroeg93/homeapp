<template>
  <ul class="divide-y divide-gray-200">
    <template v-for="lightbulb in lightbulbs" :key="lightbulb.id" class="py-4 flex">
      <LightbulbListEntry v-bind:lightbulb="lightbulb"/>
    </template>
  </ul>
</template>

<script setup lang="ts">
import {defineProps} from 'vue'
import LightbulbListEntry from "../../components/lightning/LightbulbListEntry.vue";

defineProps({
  lightbulbs: {}
})
</script>