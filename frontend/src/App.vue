<script setup lang="ts">

import {RouterView } from 'vue-router';
import Navigation from './components/navigation/Navigation.vue';
import { ref } from 'vue';
import Login from './views/Login.vue';

const auth = ref(false);

const switchAuth = () => {
  auth.value = !auth.value
}

</script>


<template>
  <div class="app">
    <div v-if="auth" class="wrapper">
      <Navigation />
      <main class="main">
        <RouterView />
      </main>
    </div>
    <Login v-else :switchAuth="switchAuth"/>
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