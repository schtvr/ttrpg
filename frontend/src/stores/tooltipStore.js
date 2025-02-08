import { defineStore } from "pinia";
import { computed, ref } from "vue";

export const useTooltipStore = defineStore("tooltip", () => {
  const tooltipKey = ref(null);

  const descriptions = {
    agility_attack: {
      title: "Agility Attack",
      description: "Rolls an attack using your agility skill modifier.",
    },
    healing_potion: {
      title: "Healing Potion",
      description: "Restores 10 HP when used.",
    },
    fireball_spell: {
      title: "Fireball",
      description: "Deals 3d6 fire damage in a 10ft radius.",
    },
  };

  const tooltipContent = computed(() => tooltipKey.value ? descriptions[tooltipKey.value] : null);

  const setTooltip = (key) => {
    tooltipKey.value = key;
  };

  const clearTooltip = () => {
    tooltipKey.value = null;
  };

  return { tooltipKey, tooltipContent, setTooltip, clearTooltip };
});
