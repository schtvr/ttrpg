<template>
  <div class="chatbox">
    <div class="messages">
      <div v-for="(msg, index) in messages" :key="index" class="message">
        <strong>{{ msg.sender }}:</strong> {{ msg.content }}
      </div>
    </div>
    <div class="input-box">
      <input v-model="newMessage" @keyup.enter="sendMessage" placeholder="Type your message here..." />
      <button @click="sendMessage">Send</button>
    </div>
  </div>
</template>

<script>
export default {
  name: "ChatBox",
  data() {
    return {
      messages: [], // List of chat messages
      newMessage: "", // The message being typed
      socket: null,  // WebSocket connection
    };
  },
  methods: {
    connectWebSocket() {
      // Create WebSocket connection
      this.socket = new WebSocket("ws://localhost:8080/ws");

      // Listen for messages from the server
      this.socket.onmessage = (event) => {
        const msg = JSON.parse(event.data);
        this.messages.push(msg);
      };

      // Log errors
      this.socket.onerror = (error) => {
        console.error("WebSocket error:", error);
      };

      // Handle WebSocket connection close
      this.socket.onclose = () => {
        console.log("WebSocket connection closed. Reconnecting...");
        setTimeout(this.connectWebSocket, 3000); // Attempt to reconnect after 3 seconds
      };
    },
    sendMessage() {
      if (this.newMessage.trim() === "") return;

      // Send message through WebSocket
      const message = {
        sender: "User", // Replace this with actual sender's name/ID
        content: this.newMessage,
      };

      this.socket.send(JSON.stringify(message));
      this.newMessage = ""; // Clear input field
    },
  },
  mounted() {
    this.connectWebSocket(); // Connect to WebSocket when the component is mounted
  },
  beforeUnmount() {
    if (this.socket) this.socket.close(); // Close WebSocket when component is destroyed
  },
};
</script>

<style scoped>
.chatbox {
  border: 1px solid #ccc;
  padding: 16px;
  max-width: 400px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.messages {
  flex: 1;
  height: 300px;
  overflow-y: auto;
  border: 1px solid #eee;
  padding: 8px;
  background-color: #f9f9f9;
}

.message {
  margin: 4px 0;
}

.input-box {
  display: flex;
  gap: 8px;
}

input {
  flex: 1;
  padding: 8px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

button {
  padding: 8px 12px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

button:hover {
  background-color: #0056b3;
}
</style>