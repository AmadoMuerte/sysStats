<script setup lang="ts">

const socket = new WebSocket('ws://localhost:8080/api/v1/websocket/metrics'); 
socket.onopen = (event) => {
  console.log("WebSocket connection opened");
};

socket.onmessage = (event) => {
  console.log("WebSocket message received:", JSON.parse(event.data));
  socket.close()
};

</script>

<template>
  
</template>