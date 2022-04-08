<script setup lang="ts">
import UserList from '@/components/UserList.vue'
</script>

<template>
  <section>
    <UserList class="user_list_component"/>
  </section>
</template>

<style scoped>
.user_list_component {
  width: 300px;
  margin-left: 2em;
}
</style>
