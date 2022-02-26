<script setup lang="ts">
import Profile from '@/components/Profile.vue'
import FriendList from '@/components/FriendList.vue'
import UserList from '@/components/UserList.vue'
</script>

<template>
  <section>
    <Profile />
    <FriendList/>
    <UserList />
  </section>
</template>

<style scoped>
section {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}
</style>
