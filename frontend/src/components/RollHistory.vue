<!-- RollHistory.vue -->
<template>
  <div class="roll-history">
    <h2>Roll History</h2>
    <ul>
      <li v-for="roll in rolls" :key="roll.id">
        {{ roll.timestamp }} - {{ roll.description }}: {{ roll.result }}
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  data() {
    return {
      rolls: [
        { id: 1, timestamp: '2024-12-22 12:00', description: 'Agility Roll', result: '2d6: [0, +1] = 1' },
      ],
    };
  },
};
</script>

<style scoped>
.roll-history {
  margin-top: 1rem;
  padding: 1rem;
  border: 1px solid #ccc;
  border-radius: 8px;
}

ul {
  list-style: none;
  padding: 0;
}

li {
  margin-bottom: 0.5rem;
}
</style>
