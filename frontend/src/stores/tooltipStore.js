import { useCharacterStore } from '@/stores/characterStore';
import { defineStore } from "pinia";
import { computed, ref } from "vue";


export const useTooltipStore = defineStore("tooltip", () => {
  const tooltipKey = ref(null);

  const definitions = {
    // skills
    agility: "skill: agility - physical dexterity the skill of moving quickly and easily",
    brawl: "skill: contacts - long-term charisma the skill of knowing and making connections with people",
    insight: "skill: insight - applied wisdom the skill of understanding a situation and all its constituent parts",
    operate: "skill: operate - applied intelligence the skill of interacting with technology",
    perform:    "skill: perform - social dexteritythe skill of entertaining",
    persuade: "skill: persuade - social strength the skill of convincing others",
    physique: "skill: physique - physical strength the skill of utilizing strength",
    shoot:    "skill: shoot - ranged combatthe skill of hitting a target",
    stealth: "skill: stealth - movement dexterity the skill of unnoticed action",

    // actions
    attack: "action: attack - initiates a volley to damage a target (e.g. punch, grapple, insult, etc.)",
    defend: "action: defend - initiates a volley to prevent damage (e.g. block, dodge, disable, etc.)",
    overcome: "action: overcome - solves an obstacle",
    support: "action: support - bolsters one or more allies",
  }

  const generateTooltip = (inputs) => {
    const characterStore = useCharacterStore()
    const tooltip = {
      terms: [],
      descriptions: [],
    }
    const actions = ["attack", "defend", "overcome", "empower"]
    let action = ""
    let skill = ""

    inputs.forEach((term) => {
      tooltip.terms.push(term)
      tooltip.descriptions.push(definitions[term])

      // find sources
      if (actions.includes(term)) {
        action = term
      } else {
        skill = term
      }

      if (characterStore.modifiers[skill][action] && characterStore.modifiers[skill][action].sources.length){
        tooltip.sources = characterStore.modifiers[skill][action].sources
      }
    })
    return tooltip
  }

  const tooltipContent = computed(() => tooltipKey.value ? generateTooltip(tooltipKey.value) : null)


  const setTooltip = (key) => {
    tooltipKey.value = key;
  };

  const clearTooltip = () => {
    tooltipKey.value = null;
  };

  return { tooltipKey, tooltipContent, setTooltip, clearTooltip };
});
