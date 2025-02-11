<template>
  <div v-if="tooltipContent" class="tooltip">
    <p>{{ tooltipContent.terms }}</p>
    <p v-for="d in tooltipContent.descriptions" :key="d">{{ d }}</p>
    <div v-if="tooltipContent.sources">
      <p>sources:</p>
      <p v-for="d in tooltipContent.sources" :key="d">
        {{ d.effects[0].modifier > -1 ? `+${d.effects[0].modifier}` : d.effects[0].modifier }}: {{ d.name }}
      </p>
    </div>
  </div>
</template>

<script>
import { computed } from "vue";
import { useTooltipStore } from "@/stores/tooltipStore";

export default {
  setup() {
    const tooltipStore = useTooltipStore();

    const tooltipContent = computed(() => tooltipStore.tooltipContent);

    return { tooltipContent };
  },
};
</script>

<style scoped>
.tooltip {
  position: absolute;
  top: 0;
  left: 100%;
  white-space: nowrap;
  background: rgba(0, 0, 0, 0.8);
  color: white;
  padding: 8px;
  border-radius: 5px;
  box-shadow: 2px 2px 10px rgba(0, 0, 0, 0.3);
  z-index: 1000;
}
</style>