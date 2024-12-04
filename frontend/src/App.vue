<script setup lang="ts">

// @ts-ignore
import Footer from './components/Footer.vue';
import { RouterLink, RouterView } from 'vue-router';
import Navigation from './components/navigation/Navigation.vue';

// const socket = new WebSocket('ws://localhost:8080/api/v1/websocket/metrics'); 
// socket.onopen = (event) => {
//   console.log("WebSocket connection opened");
// };

// socket.onmessage = (event) => {
//   console.log("WebSocket message received:", JSON.parse(event.data));
//   socket.close()
// };

</script>


<template>
  <div class="app">
    <div class="wrapper">
      <Navigation />
      <main class="main">
        <RouterView />
      </main>
    </div>
  </div>
</template>

<style>
.app {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  min-height: 100vh;
}

.wrapper {
  display: flex;
}

.main {
  padding: 30px 0;
  margin-left: 90px;
}

</style>