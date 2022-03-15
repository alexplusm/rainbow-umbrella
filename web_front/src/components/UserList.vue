<script setup lang="ts">
import { ref } from "vue";
import { useUserStore } from "@/stores/user";
import ProfileMini from "@/components/ProfileMini.vue";

const userStore = useUserStore();
const searchControl = ref("");

function search() {
  console.log("search: ", searchControl.value);
}
</script>

<template>
  <section class="wrap">
    <div class="text-h6 q-pb-sm">Users</div>

    <div class="q-pb-sm row">
      <q-input
        class="col"
        label="Search user"
        v-model="searchControl"
        v-on:change="search"
        clearable
      />

      <q-btn
        flat
        color="deep-orange"
        icon="search"
      />
    </div>

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
