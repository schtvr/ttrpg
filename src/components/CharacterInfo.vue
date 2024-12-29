<!-- CharacterInfo.vue -->
<template>

  <div class="character__container">
    <div class="character__basic-info">
      <h2>{{ character.name }}</h2>
      <h3>{{ character.epithet }}</h3>
    </div>
    <div class="character__icons">
      <p>icon</p>
    </div>
    <div class="character__status">
      <h3>momentum</h3>
      <h3>wounds</h3>
    </div>
    <div class="character__tabs">
      <button v-for="tab in tabs"
        :key="tab.name"
        :class="{ active: activeTab === tab.name }"
        @click="activeTab = tab.name">
        {{ tab.label }}
      </button>
    </div>

    <div class="character__tab-content">
      <div v-if="activeTab === 'about'" >
        <h3>about</h3>
        <p>{{character.about}}</p>
      </div>
      <InventoryTable v-if="activeTab === 'wiki'" />
    </div>
  </div>
</template>

<script>
import InventoryTable from "./InventoryTable.vue"

export default {
  components: {
    InventoryTable
  },
  data() {
    return {
      character: {
        name: 'joey',
        epithet: 'bingo',
        health: 100,
      },
      tabs: [
        { name: 'about', label: 'about' },
        { name: 'inventory', label: 'inventory' },
        { name: 'skills', label: 'skills' },
        { name: 'wounds', label: 'wounds' },
      ],
      activeTab: 'about',
    };
  },
};
</script>

<style scoped>
.character__container {
  padding: 1rem;
  display: grid;
  grid-template-columns: 2fr .75fr .75fr 4fr;
  grid-template-rows: repeat(2, 1fr);
  grid-column-gap: 10px;
  grid-row-gap: 10px;
}

.character__status { grid-area: 2 / 1 / 2 / 2; }
.character__tabs { grid-area: 1 / 3 / 2 / 3; }
.character__tab-content { grid-area: 1 / 4 / 2 / 4; }
.character__basic-info { grid-area: 1 / 1 / 1 / 1; }
.character__icons { grid-area: 1 / 2 / 1 / 2; }


</style>
