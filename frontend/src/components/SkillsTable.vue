<!-- SkillsTable.vue -->
<template>
  <div class="skills">
    <div class="tooltip">This is the tooltip text!</div>
    <div v-if="tooltipVisible" class="tooltip">This is the hidden tooltip text!</div>
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
            <button class="tooltip-container" @mouseenter="(e) => showTooltip(e, skill, action)"
              @mouseleave="hideTooltip" @click="roll(skill[action].modifier, 4)">
              {{ skill[action].modifier < 0 ? skill[action].modifier : `+${skill[action].modifier}` }} </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
import { useCharacterStore } from '@/stores/characterStore'
import { sendMessage } from "@/services/websocketService";

export default {
  setup() {
    const characterStore = useCharacterStore()
    characterStore.computeModifiers(characterStore)

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
      roll
    }
  },
}
</script>

<style scoped>
.skills {
  margin-top: 1rem;
  position: relative;
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

.tooltip {
  position: absolute;
  bottom: 88%;
  /* Position above the container */
  left: 107%;
  background-color: #333;
  color: white;
  padding: 5px 10px;
  border-radius: 4px;
  font-size: 14px;
  white-space: nowrap;
  z-index: 1000;
  opacity: 0;
  visibility: hidden;
  transition:
    opacity 0.2s ease,
    visibility 0.2s ease;
}

.tooltip-container:hover .tooltip {
  opacity: 1;
  visibility: visible;
}
</style>