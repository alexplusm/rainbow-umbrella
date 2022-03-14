<script setup lang="ts">
import { ref } from "vue";
import { useUserStore } from "@/stores/user";
import ProfileMini from "@/components/ProfileMini.vue";

const userStore = useUserStore();
const searchControl = ref("");
</script>

<template>
  <section>
    <div class="text-h6">Users</div>

    <q-input v-model="searchControl" label="Search user" />

    <q-list class="user_list" bordered separator>
      <ProfileMini
          v-for="user in userStore.usersList"
          :key="user.id"
          :user="user"
          :show-actions="true"
      />
    </q-list>
  </section>
</template>

<style scoped>
section {
  display: flex;
  flex-direction: column;
}

.user_list {
  overflow: scroll;
}
</style>
