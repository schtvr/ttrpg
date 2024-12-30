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
          <td><button>{{ skill.attack < 0 ? skill.attack : `+${skill.attack}` }}</button></td>
          <td><button>{{ skill.defend < 0 ? skill.defend : `+${skill.defend}` }}</button></td>
          <td><button>{{ skill.overcome < 0 ? skill.overcome : `+${skill.overcome}` }}</button></td>
          <td><button>{{ skill.empower < 0 ? skill.empower : `+${skill.empower}` }}</button></td>
        </tr>
      </tbody>
    </table>
  </div>
</template>
<script>
import { useCharacterStore } from '@/stores/characterStore';

export default {
  setup() {
    const characterStore = useCharacterStore();
    characterStore.computeModifiers(characterStore);

    // Example of accessing state and actions
    const addItem = () => {
      characterStore.addToInventory({ name: 'Health Potion', quantity: 1 });
    };

    return {
      modifiers: characterStore.modifiers,
      characterName: characterStore.characterName,
      addItem,
    };
  },
};
</script>
<!-- <script>
export default {
  data() {
    return {
      skills: [
        { name: 'Agility', attack: '+1', defend: '+2', overcome: '+1', empower: '-1' },
        { name: 'Brawl', attack: '+4', defend: '+2', overcome: '+0', empower: '-1' },
        { name: 'Contacts', attack: '-1', defend: '+1', overcome: '+2', empower: '+4' },
        { name: 'Shoot', attack: '+2', defend: '+1', overcome: '+0', empower: '-1' },
      ],
    };
  },
};
</script> -->

<style scoped>
.skills {
  margin-top: 1rem;
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

