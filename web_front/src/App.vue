<script setup lang="ts">
import { RouterView } from 'vue-router'
</script>

<template>
  <header>
    <div class="logo">
      🦄 & 🌈
    </div>
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
  height: 46px;
}

.logo {
  display: block;
  font-size: 24px;
}

.main {
  justify-content: center;
}
</style>
