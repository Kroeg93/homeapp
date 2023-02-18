<template>
  <div class="container mx-auto">
    <h1>Beleuchtung</h1>

    <ul>
      <li v-for="(lightbulb, index) in state.lightbulbs" :key="lightbulb.id">
        <div class="ml-3">
          <p class="text-sm font-medium text-gray-900">{{ lightbulb.id }}</p>
          <p class="text-sm text-gray-500">{{ lightbulb.name }}</p>
        </div>
      </li>
    </ul>

  </div>
</template>

<script>
  import { reactive, onMounted } from 'vue'
  import {fetchLightbulbs} from '../../api/lightning'

  export default {
    setup() {
      const state = reactive({
        lightbulbs: null
      })

      onMounted(async () => {state.lightbulbs = await fetchLightbulbs()})

      console.log(state.lightbulbs)

      return {
        state
      }
    }
  }
</script>

