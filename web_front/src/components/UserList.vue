<script setup lang="ts">
import { ref } from "vue";
import { useUserStore } from "@/stores/user";
import ProfileMini from "@/components/ProfileMini.vue";

const userStore = useUserStore();
const searchControl = ref("");
</script>

<template>
  <section class="wrap">
    <div class="text-h6 q-pb-sm">Users</div>

    <q-input class="q-pb-sm" label="Search user" v-model="searchControl" />

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
.wrap {
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
}

.user_list {
  max-height: 70vh;
  overflow: scroll;
}
</style>
