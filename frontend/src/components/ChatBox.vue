<template>
  <div class="chatbox-container">
    <!-- Messages Section -->
    <div class="messages">
      <div v-for="(message, index) in messages" :key="index" class="message">
        <span class="timestamp">({{ message.datetime }}) </span>
        <strong>{{ message.sender }}: </strong>
        {{ message.content }}
      </div>
    </div>

    <!-- Input Section -->
    <div class="input-container">
      <input v-model="newMessage" @keyup.enter="sendMessage" placeholder="Type your message..." class="message-input" />
      <button @click="sendMessage" class="send-button">Send</button>
    </div>
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
        const now = new Date();
        const formattedTime = now.toTimeString().slice(0, 5); // Format HH:MM

        sendMessage({
          sender: "steve", // Replace with dynamic sender name if needed
          content: newMessage.value,
          datetime: formattedTime,
        });

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
.chatbox-container {
  display: flex;
  flex-direction: column;
  height: 100%;
  width: 100%;
  border: 1px solid #ccc;
  border-radius: 8px;
  overflow: hidden;
}

/* Messages Section */
.messages {
  flex: 1;
  overflow-y: auto;
  padding: 16px;
  /* background-color: #f9f9f9; */
}

.message {
  margin-bottom: 10px;
}

.timestamp {
  color: gray;
  font-size: 0.8em;
}

/* Input Section */
.input-container {
  display: flex;
  align-items: center;
  padding: 8px;
  border-top: 1px solid #ccc;
  /* background-color: #fff; */
}

.message-input {
  flex: 1;
  padding: 10px;
  font-size: 1rem;
  /* border: 1px solid #ccc; */
  border-radius: 4px;
  outline: none;
}

.message-input:focus {
  border-color: #007bff;
  box-shadow: 0 0 3px rgba(0, 123, 255, 0.5);
}

.send-button {
  margin-left: 8px;
  padding: 10px 16px;
  font-size: 1rem;
  background-color: #007bff;
  /* color: #fff; */
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

.send-button:hover {
  background-color: #0056b3;
}
</style>

setup() {
connectWebSocket("ws://localhost:8080/ws")