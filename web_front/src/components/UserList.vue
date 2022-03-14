<script setup lang="ts">
import { ref } from "vue";
import { useUserStore } from "@/stores/user";
import ProfileMini from "@/components/ProfileMiniOld.vue";

const userStore = useUserStore();
const searchControl = ref("");
</script>

<template>
  <section>
    <q-card flat bordered>
      <q-card-section>
        <div class="text-h6">Users</div>
      </q-card-section>

      <q-card-section class="q-pt-none">
        <q-input v-model="searchControl" label="Search user" />
      </q-card-section>

      <q-card-section class="scroll overflow-hidden" style="height: 60vh">
        <ProfileMini
            v-for="user in userStore.usersList"
            :key="user.id"
            :user="user"
            :show-actions="true"
        />
      </q-card-section>
    </q-card>
  </section>
</template>

<style scoped>
</style>
