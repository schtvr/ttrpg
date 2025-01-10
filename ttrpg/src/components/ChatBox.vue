<!-- ChatBox.vue -->
<template>
  <div class="chat-box">
    <h2>Chat</h2>
    <div class="messages">
      <div v-for="message in messages" :key="message.id" class="message">
        <strong>{{ message.user }}:</strong> {{ message.text }}
      </div>
    </div>
    <input v-model="newMessage" @keyup.enter="sendMessage" placeholder="Type your message..." />
    <button @click="sendMessage">Send</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      messages: [],
      newMessage: '',
      socket: null,
    };
  },
  methods: {
    initializeSocket() {
      // Establish connection with the server
      this.socket = new WebSocket('ws://your-server-url'); // Replace with your WebSocket server URL

      // Handle incoming messages
      this.socket.onmessage = (event) => {
        const message = JSON.parse(event.data);
        this.messages.push({
          id: message.id,
          user: message.user,
          text: message.text,
        });
      };

      // Handle connection close
      this.socket.onclose = () => {
        console.log('Connection closed.');
      };

      // Handle errors
      this.socket.onerror = (error) => {
        console.error('WebSocket error:', error);
      };
    },
    sendMessage() {
      if (this.newMessage.trim() !== '') {
        const message = {
          id: Date.now(),
          user: 'Player', // Replace with dynamic user identification if needed
          text: this.newMessage,
        };

        // Send the message to the server
        this.socket.send(JSON.stringify(message));

        // Optionally add the message locally (if the server doesn't echo it back)
        this.messages.push(message);

        this.newMessage = '';
      }
    },
  },
  mounted() {
    this.initializeSocket();
  },
  beforeUnmount() {
    if (this.socket) {
      this.socket.close();
    }
  },
};
</script>

<style scoped>
.chat-box {
  margin-top: 1rem;
  padding: 1rem;
  height: 100%;
  border: 1px solid #ccc;
  border-radius: 8px;
}

.messages {
  max-height: 200px;
  overflow-y: auto;
  margin-bottom: 1rem;
}

.message {
  margin-bottom: 0.5rem;
}

input {
  width: calc(100% - 60px);
  padding: 0.5rem;
  margin-right: 0.5rem;
}

button {
  padding: 0.5rem;
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

<!-- Explanation -->
<!--
This component connects to a WebSocket server to handle real-time communication. Here’s the breakdown:

1. **WebSocket Initialization**:
   - When the component is mounted, it establishes a WebSocket connection using the `initializeSocket` method.
   - The `ws://your-server-url` should be replaced with the actual WebSocket server address.

2. **Handling Incoming Messages**:
   - The `onmessage` event listens for new messages from the server and adds them to the `messages` array, which updates the chat view in real-time.

3. **Sending Messages**:
   - The `sendMessage` method sends the `newMessage` to the server as a JSON string.
   - It also appends the message to the local `messages` array for immediate feedback, assuming the server might not echo it back.

4. **Connection Management**:
   - The `onclose` and `onerror` events handle connection closures and errors, providing logging for debugging.

5. **Component Cleanup**:
   - When the component is unmounted, the WebSocket connection is gracefully closed to prevent memory leaks.

Replace `your-server-url` with your actual WebSocket endpoint. Ensure the server is set up to broadcast messages to connected clients.
-->
