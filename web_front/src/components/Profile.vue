<template>
  <section>
    KEKW
<!--    {{ user.login }}-->
  </section>
</template>

<script lang="ts" setup>
import router from '@/router'
import {UserM} from "@/models/user";
import {onUnmounted} from "vue";
import {useUserStore} from "@/stores/counter";


const login = router.currentRoute.value.params['login'];
console.log("curr login", login);

const sessionId = localStorage.getItem("X-SessionId");

console.log("sessionId", sessionId)

// withAsyncContext()

const userStore = useUserStore();
userStore.retrieve(login as string);

</script>

<style scoped>
</style>