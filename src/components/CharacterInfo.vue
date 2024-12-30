<!-- CharacterInfo.vue -->
<template>

  <div class="character__container">
    <div class="character__basic-info">
      <h2>{{ character.name }}</h2>
      <h3>{{ character.epithet }}</h3>
      <div class="character__status">
        <h3>momentum [ ] [ ] [ ] [ ]</h3>
        <h3>wounds   [ ] [ ] [ ] [ ]</h3>
      </div>
    </div>

    <div class="character__tab-content">
      <div class="character__tabs">
        <button v-for="tab in tabs"
          :key="tab.name"
          :class="{ active: activeTab === tab.name }"
          @click="activeTab = tab.name">
          {{ tab.label }}
        </button>
      </div>
      <div v-if="activeTab === 'about'" >
        <h3>about</h3>
        <p>{{character.about}}</p>
      </div>
      <InventoryTable v-if="activeTab === 'inventory'" />
    </div>
  </div>
</template>

<script>
import InventoryTable from "./InventoryTable.vue"
import { useCharacterStore } from '@/stores/characterStore';

export default {
  components: {
    InventoryTable
  },
  setup() {
    const characterStore = useCharacterStore();

    return {
      character: {
        name: characterStore.name,
        epithet: characterStore.epithet,
        about: characterStore.about,
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
  display: flex;
  height: 100%;
  gap: 10px;
}

.character__tabs {
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
  gap: 3px;
}

.character__tabs > button {
  flex-grow: 1;
}
.character__tab-content {
  display: flex;
  flex-grow: 1;
  flex-direction: column;
  align-items: stretch;
}


</style>
