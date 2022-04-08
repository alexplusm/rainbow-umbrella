<script setup lang="ts">
import Profile from '@/components/Profile.vue'
import FriendList from '@/components/FriendList.vue'
</script>

<template>
  <section>
    <Profile class="profile_component" />
    <FriendList class="friend_list_component"/>
  </section>
</template>

<style scoped>
section {
  width: 100%;
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
}

.profile_component {
  width: 250px;
}

.friend_list_component {
  width: 300px;
  margin-left: 2em;
}
</style>
