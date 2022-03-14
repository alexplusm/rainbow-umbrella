<script setup lang="ts">
import ProfileMini from "@/components/ProfileMini.vue";
import { useUserStore } from "@/stores/user";

const userStore = useUserStore();
</script>

<template>
  <section class="q-pa-md" style="max-width: 350px">
    <div class="text-h6">Friend List</div>

    <q-list bordered separator>
      <ProfileMini
        class="profile_mini"
        :user="userStore.currentUser"
        :showActions="true"
      />

      <ProfileMini
        class="profile_mini"
        :user="userStore.currentUser"
        :showActions="true"
      />

      <ProfileMini
        class="profile_mini"
        :user="userStore.currentUser"
        :showActions="true"
      />
    </q-list>
  </section>
</template>

<style scoped>
section {
  display: flex;
  flex-direction: column;
}
</style>
