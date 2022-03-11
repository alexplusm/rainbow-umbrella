<script setup lang="ts">
import ProfileMini from "@/components/ProfileMini.vue"
import {useUserStore} from "@/stores/user";

const userStore = useUserStore();


</script>

<template>
  <section>
    FriendList

    <ProfileMini :user="userStore.currentUser" :show-actions="false"/>
  </section>
</template>

<style scoped>
</style>