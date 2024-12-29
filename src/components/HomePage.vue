<!-- Home.vue -->
<template>
  <div class="home">
    <h1>Welcome to the TTRPG App</h1>
    <div class="sections">
      <CharacterInfo />
      <Skills />
      <RollHistory />
      <ChatBox />
    </div>
  </div>
</template>

<script>
import CharacterInfo from "./components/CharacterInfo.vue";
import Skills from "./components/Skills.vue";
import RollHistory from "./components/RollHistory.vue";
import ChatBox from "./components/ChatBox.vue";

export default {
  components: {
    CharacterInfo,
    Skills,
    RollHistory,
    ChatBox,
  },
};
</script>

<style scoped>
.home {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 1rem;
}

.sections {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
  gap: 1rem;
  width: 100%;
  max-width: 1200px;
}
</style>
