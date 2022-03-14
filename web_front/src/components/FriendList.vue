<script setup lang="ts">
import ProfileMini from "@/components/ProfileMini.vue";
import { useUserStore } from "@/stores/user";

const userStore = useUserStore();
</script>

<template>
  <section class="wrap">
    <div class="text-h6 q-pb-sm">Friend List</div>

    <q-list class="user_list" bordered separator>
      <ProfileMini
        class="profile_mini"
        :user="userStore.currentUser"
        :showActions="false"
      />

      <ProfileMini
        class="profile_mini"
        :user="userStore.currentUser"
        :showActions="false"
      />

      <ProfileMini
        class="profile_mini"
        :user="userStore.currentUser"
        :showActions="false"
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
