<template>
  <div>
    <h2>Chat</h2>
    <div class="messages">
      <div v-for="(message, index) in messages" :key="index">
        {{ message.content }}
      </div>
    </div>
    <input v-model="newMessage" @keyup.enter="sendMessage" placeholder="Type a message..." />
    <button @click="sendMessage">Send</button>
  </div>
</template>

<script>
import { ref } from "vue";
import { connectWebSocket, useWebSocketMessages, sendMessage } from "../services/websocketService.js";

export default {
  setup() {
    connectWebSocket("ws://localhost:8080/ws")
    const messages = useWebSocketMessages(); // Shared state
    const newMessage = ref("");

    const sendMessageHandler = () => {
      if (newMessage.value.trim()) {
        sendMessage({ content: newMessage.value }); // Send message to the server
        newMessage.value = ""; // Clear input
      }
    };

    return {
      messages,
      newMessage,
      sendMessage: sendMessageHandler,
    };
  },

};
</script>

<style scoped>
.messages {
  height: 300px;
  overflow-y: auto;
  border: 1px solid #ccc;
  padding: 10px;
}
</style>