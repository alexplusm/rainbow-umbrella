<template>
  <section>
    KEKW
    {{ userStore.currentUser?.login }}
  </section>
</template>

<script lang="ts" setup>
import router from '@/router'
import {useUserStore} from "@/stores/user";

const userStore = useUserStore()

const login = router.currentRoute.value.params['login'];
console.log("curr login", login);

const sessionId = localStorage.getItem("X-SessionId");

console.log("sessionId", sessionId);

userStore.retrieve(login as string);
</script>

<style scoped>
</style>