<template>
  <div class="min-w-full">
    <img class="h-10 w-10 rounded-full" alt=""/>
    <p class="text-sm font-medium text-gray-900">{{ lightbulb.name }}</p>
    <p class="text-sm text-gray-500">{{ lightbulb.type }} | {{ lightbulb.productname }}</p>
    <div class="content-end">
      <v-btn>
        Button
      </v-btn>
    </div>
  </div>
</template>

<script setup lang="ts">
import {defineProps} from "vue";

defineProps({
  lightbulb: {}
})

</script>