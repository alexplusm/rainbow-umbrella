<template>
  <section>
    <label>
      Search User
      <input type="search" placeholder="Type ...">
    </label>

    UserList

<!--    {{ userStore.users?.length }}-->
    a {{ userStore.usersList?.length }} b
  </section>
</template>

<script setup lang="ts">
import {useUserStore} from "@/stores/user";

console.log("INIT COMPONENT!!!");

const userStore = useUserStore();

console.log("users", userStore.usersList);
</script>


<style scoped>
</style>