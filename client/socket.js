export const socket = new WebSocket("ws://localhost:8080/ws");

// WebSocket connection open event
socket.onopen = () => {
    socket.send("Hello, server!"); 
};

// WebSocket message event
socket.onmessage = (event) => {
    console.log("Message from server:", event.data);
};

// WebSocket connection close event
socket.onclose = () => {
    // contentId.append("WebSocket connection closed");
    console.log("WebSocket connection closed");
};

// WebSocket error event
socket.onerror = (error) => {
    console.log("WebSocket error:", error);
};

