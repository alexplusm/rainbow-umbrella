<template>
  <section>
    <label>
      Search User
      <input type="search" placeholder="Type ...">
    </label>

    UserList

    {{ userStore.users?.length }}
  </section>
</template>

<script setup lang="ts">
import {useUserStore} from "@/stores/user";

console.log("INIT COMPONENT!!!");

const userStore = useUserStore();
// const users = userStore.users;

// TODO: DON`T WORK
console.log("users", userStore.users);

</script>


<style scoped>
</style>