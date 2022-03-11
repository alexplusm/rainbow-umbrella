<script setup lang="ts">
import ProfileMini from "@/components/ProfileMini.vue";
import { useUserStore } from "@/stores/user";

const userStore = useUserStore();
</script>

<template>
  <section>
    <q-card flat bordered>
      <q-card-section>
        <div class="text-h6">Friend List</div>
      </q-card-section>

      <q-card-section>
        <ProfileMini :user="userStore.currentUser" :show-actions="false"/>
        <ProfileMini :user="userStore.currentUser" :show-actions="false"/>
        <ProfileMini :user="userStore.currentUser" :show-actions="false"/>
      </q-card-section>
    </q-card>
  </section>
</template>

<style scoped>
section {
  display: flex;
  flex-direction: column;
}
</style>
