<template>
  <div class="grid grid-cols-3 gap-4 p-4">
    <button v-for="btn in aspects" :key="btn" class="px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-700">
      {{ btn.name }}
    </button>
  </div>
  <div class="flex justify-center mt-4">
    <button @click="generateColumns" class="px-4 py-2 bg-green-500 text-white rounded-lg hover:bg-green-700">
      Regenerate Buttons
    </button>
  </div>
</template>

<script>
import { ref, onMounted } from 'vue';
import { useCharacterStore } from '@/stores/characterStore'

export default {
  setup() {
    const aspects = useCharacterStore().aspects
    return {
      aspects
    }

  }
}


const columns = ref([[], [], []]);

const generateColumns = () => {
  columns.value = columns.value.map(() => Array.from({ length: Math.floor(Math.random() * 6) + 1 }, (_, i) => i + 1));
};

onMounted(generateColumns);
</script>

<style scoped>
.grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
}
</style>