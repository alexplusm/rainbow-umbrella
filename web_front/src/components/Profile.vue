<script lang="ts" setup>
import ProfileFriendshipStatus from "@/components/ProfileFriendshipStatus.vue";
import { useUserStore } from "@/stores/user";

const userStore = useUserStore();
</script>

<template>
  <q-card flat bordered>
    <q-item>
      <q-item-section avatar>
        <q-avatar>
          <img alt="avatar" :src="userStore.currentUser.avatarUrl">
        </q-avatar>
      </q-item-section>

      <q-item-section>
        <q-item-label>
          {{ userStore.currentUser.firstName }} {{ userStore.currentUser.lastName }}
        </q-item-label>
        <q-item-label caption>{{ userStore.currentUser.login }}</q-item-label>
      </q-item-section>

      <ProfileFriendshipStatus
          :userId="userStore.currentUser.id"
          :friendshipStatus="userStore.currentUser.friendshipStatus"
      />
    </q-item>

    <q-separator />

    <q-card-section>
      <ul>
        <li>Age: {{ userStore.currentUser.age }}</li>
        <li>City: {{ userStore.currentUser.city}}</li>
        <li>Gender: {{ userStore.currentUser.gender }}</li>
      </ul>
    </q-card-section>

    <q-separator />

    <q-card-section>
      <div class="text-h6">Interests</div>

      <q-badge
        class="q-ml-sm q-mb-sm"
        rounded color="green"
        v-for="interest in userStore.currentUser.interests"
        :label="interest"
      />
    </q-card-section>
  </q-card>
</template>

<style scoped>
</style>
