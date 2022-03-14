<script setup lang="ts">
import { RouterView } from 'vue-router';
import { useAppStore } from "@/stores/app";

const appStore = useAppStore();
</script>

<template>
  <header>
    <router-link class="logo" :to="appStore.logoNavLink">
      <q-btn size="0.8em" flat style="color: #FF0080" label="🦄 & 🌈" />
    </router-link>

  </header>
  <RouterView class="main"/>
</template>

<style>

#app {
  display: flex;
  flex-direction: column;
  max-width: 1280px;
  margin: 0 auto;
  padding: 1rem;

  font-weight: normal;
}

header {
  width: 100vw;
  height: 8vh;
}

.logo {
  display: block;
  font-size: 24px;
  text-decoration: none;
  color: black;
}

.main {
  justify-content: center;
  height: 85vh;
}
</style>
