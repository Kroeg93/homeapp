<template>
  <div class="container mx-auto">
    <h1>Beleuchtung</h1>

    <LightbulbList v-bind:lightbulbs="lightbulbs" ></LightbulbList>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
  import {fetchLightbulbs} from '../../api/lightning'
  import LightbulbList from "../../components/lightning/LightbulbList.vue";

  const lightbulbs = ref()

  const getLightbulbs = async () => {
    lightbulbs.value = await fetchLightbulbs()
  }

  onMounted(async () => getLightbulbs())
</script>

