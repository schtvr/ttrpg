<!-- SkillsTable.vue -->
<template>
  <div class="skills">
    <h2>Skills</h2>
    <table>
      <thead>
        <tr>
          <th>Skill</th>
          <th>Attack</th>
          <th>Defend</th>
          <th>Overcome</th>
          <th>Empower</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="skill in modifiers" :key="skill.name">
          <td>{{ skill.name }}</td>
          <td v-for="action in ['attack', 'defend', 'overcome', 'empower']" :key="action">
            <button @mouseenter="tooltipStore.setTooltip([skill.name, action])" @mouseleave="tooltipStore.clearTooltip"
              @click="roll(skill[action].modifier, 4)">
              {{ skill[action].modifier < 0 ? skill[action].modifier : `+${skill[action].modifier}` }} </button>
          </td>
        </tr>
      </tbody>
    </table>
    <div class="tooltip-container">
      <ToolTip />
    </div>

  </div>
</template>

<script>
import { useCharacterStore } from '@/stores/characterStore'
import { useTooltipStore } from "@/stores/tooltipStore";
import ToolTip from "@/components/ToolTip.vue";
import { sendMessage } from "@/services/websocketService";

export default {
  components: { ToolTip },
  setup() {
    const characterStore = useCharacterStore()
    characterStore.computeModifiers(characterStore)

    const tooltipStore = useTooltipStore();


    const sendRollResults = (message) => {
      sendMessage({ sender: "steve", content: message });
      console.log("sending:", message)
    }

    const roll = (modifier, diceCount) => {
      // TODO make better for tracking roll histories
      let msg = "modifier: " + modifier
      let sum = modifier
      for (let i = 0; i < diceCount; i++) {
        const val = Math.floor(Math.random() * 3) - 1
        msg += `, roll ${i + 1}: ${val}`
        sum += val
      }
      msg += `, total roll: ${sum}`
      sendRollResults(msg)
    }

    return {
      modifiers: characterStore.modifiers,
      characterName: characterStore.characterName,
      roll,
      tooltipStore,
    }
  },
}
</script>

<style scoped>
.skills {
  margin-top: 1rem;
  position: relative;
}

.tooltip-container {
  position: relative;
  display: inline-block;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th,
td {
  border: 1px solid #ccc;
  padding: 0.5rem;
  text-align: center;
}

button {
  padding: 0.5rem;
  border: none;
  border-radius: 4px;
  background-color: #007bff;
  color: white;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}
</style>