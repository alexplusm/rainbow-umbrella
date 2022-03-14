<script setup lang="ts">
import Profile from '@/components/Profile.vue'
import FriendList from '@/components/FriendList.vue'
import UserList from '@/components/UserList.vue'
</script>

<template>
  <section>
    <Profile class="profile" />
    <FriendList class="friend_list"/>
    <UserList class="user_list_comp"/>
  </section>
</template>

<style scoped>
section {
  width: 100%;
  display: flex;
  flex-direction: row;
  justify-content: flex-start;
}

.profile {
  height: 300px;
  width: 250px;
}

.friend_list {
  width: 300px;
  margin-left: 2em;
}
.user_list_comp {
  width: 300px;
  margin-left: 2em;
}
</style>
