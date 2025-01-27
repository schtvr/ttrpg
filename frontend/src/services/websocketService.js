// websocketService.js
import { ref } from "vue";

const socket = ref(null);

export function connectWebSocket(url) {
  if (!socket.value || socket.value.readyState === WebSocket.CLOSED) {
    socket.value = new WebSocket(url);

    socket.value.onopen = () => {
      console.log("WebSocket connected");
    };

    socket.value.onerror = (error) => {
      console.error("WebSocket error:", error);
    };

    socket.value.onclose = () => {
      console.log("WebSocket connection closed");
    };
  }
}

export function sendMessage(message) {
  if (socket.value && socket.value.readyState === WebSocket.OPEN) {
    socket.value.send(JSON.stringify(message));
  } else {
    console.error("WebSocket is not open. Cannot send message:", message);
  }
}

const messages = ref([]); // Shared reactive state
let isListenerAttached = false;

export function useWebSocketMessages() {
  // Attach the WebSocket listener only once
  if (socket.value && !isListenerAttached) {
    socket.value.onmessage = (event) => {
      console.log("event:", event.data);
      const data = JSON.parse(event.data);
      messages.value.push(data); // Update shared state
    };
    isListenerAttached = true;
  } else if (!socket.value) {
    console.error("WebSocket is not initialized");
  }

  return messages; // Always return the shared reactive state
}
