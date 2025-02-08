<template>
  <div v-if="tooltipContent" class="tooltip" :style="tooltipStyles">
    <strong>{{ tooltipContent.title }}</strong>
    <p>{{ tooltipContent.description }}</p>
  </div>
</template>

<script>
import { computed, ref } from "vue";
import { useTooltipStore } from "@/stores/tooltipStore";

export default {
  setup() {
    const tooltipStore = useTooltipStore();
    const tooltipX = ref(0);
    const tooltipY = ref(0);

    const tooltipContent = computed(() => tooltipStore.tooltipContent);

    // Compute styles for tooltip positioning
    const tooltipStyles = computed(() => ({
      top: `${tooltipY.value}px`,
      left: `${tooltipX.value}px`
    }));

    return { tooltipContent, tooltipStyles };
  },
};
</script>

<style scoped>
.tooltip {
  position: fixed;
  background-color: rgba(0, 0, 0, 0.8);
  color: white;
  padding: 8px;
  border-radius: 5px;
  font-size: 0.9rem;
  pointer-events: none;
  white-space: nowrap;
  z-index: 1000;
}
</style>