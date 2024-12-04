<script setup lang="ts">

// @ts-ignore
import Footer from './components/Footer.vue';
import { RouterLink, RouterView } from 'vue-router';
import Navigation from './components/navigation/Navigation.vue';

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
  width: 100%;
  padding: 30px 30px 30px 0;
  margin-left: 90px;
}

</style>