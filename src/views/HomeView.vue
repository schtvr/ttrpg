<!-- Home.vue -->
<template>
  <div class="home">
    <div class="parent">
      <div class="sections__character">
        <CharacterInfo />
      </div>
      <div class="sections__skills-table">
        <Skills />
      </div>
      <div class="sections__omnibox">
        <div class="tabs-section">
          <div class="tabs">
            <button v-for="tab in tabs" :key="tab.name" :class="{ active: activeTab === tab.name }"
              @click="activeTab = tab.name">
              {{ tab.label }}
            </button>
          </div>

          <div class="tab-content">
            <ChatBox v-if="activeTab === 'chat'" />
            <WikiBox v-if="activeTab === 'wiki'" />
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import CharacterInfo from "../components/CharacterInfo.vue";
import ChatBox from "../components/ChatBox.vue";
import Skills from "../components/SkillsTable.vue";
import WikiBox from "../components/WikiBox.vue";

export default {
  components: {
    CharacterInfo,
    Skills,
    ChatBox,
    WikiBox,
  },
  data() {
    return {
      tabs: [
        { name: 'chat', label: 'Chat' },
        { name: 'wiki', label: 'Wiki' },
      ],
      activeTab: 'chat',
    };
  },
};
</script>

<style scoped>
.home {
  display: flexbox;
  height: 95%;
  width: 95vw;
}

.parent {
  display: grid;
  height: 95vh;
  padding: 1rem;
  grid-template-columns: 1fr 3fr;
  grid-template-rows: 20% 20% 60%;
  grid-column-gap: 12px;
  grid-row-gap: 12px;
}

.parent>div {
  border: 1px solid #000;
  border-radius: 6px;
}


.tabs-section {
  width: 100%;
  max-width: 1200px;
  border: 1px solid #ccc;
  border-radius: 8px;
  overflow: hidden;
}

.tabs {
  display: flex;
  background-color: #f4f4f4;
  border-bottom: 1px solid #ccc;
}

.tabs button {
  flex: 1;
  padding: .25rem;
  border: none;
  background: none;
  cursor: pointer;
  text-align: center;
  font-size: 1rem;
}

.tabs button.active {
  background-color: #007bff;
  color: white;
}

.tab-content {
  padding: 1rem;
}


.sections__character {
  grid-area: 1 / 1 / 1 / 3;
}

.sections__specials {
  grid-area: 2 / 1 / 2 / 1;
}

.sections__skills-table {
  grid-area: 3 / 1 / 3 / 1;
  padding: .5rem;
}

.sections__omnibox {
  grid-area: 2 / 2 / 4 / 2;
}
</style>