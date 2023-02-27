<template>
  <ul class="divide-y divide-gray-200">
    <li v-for="lightbulb in lightbulbs" :key="lightbulb.id" class="py-4 flex">
      <img class="h-10 w-10 rounded-full" alt=""/>
      <div class="min-w-full">
        <p class="text-sm font-medium text-gray-900">{{ lightbulb.name }}</p>
        <p class="text-sm text-gray-500">{{ lightbulb.type }} | {{lightbulb.productname}}</p>
        <div class="content-end">
        </div>
      </div>
    </li>
  </ul>
</template>

<script setup lang="ts">
    import {defineProps} from 'vue'
    defineProps({
      lightbulbs: {}
    })
</script>