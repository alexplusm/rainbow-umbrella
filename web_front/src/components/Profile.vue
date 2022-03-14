<script lang="ts" setup>
import ProfileFriendshipStatus from "@/components/ProfileFriendshipStatus.vue";
import { useUserStore } from "@/stores/user";

const userStore = useUserStore();
const currentUser = userStore.currentUser;
</script>

<template>
  <q-card flat bordered>
    <q-item>
      <q-item-section avatar>
        <q-avatar>
          <img alt="avatar" :src="currentUser.avatarUrl">
        </q-avatar>
      </q-item-section>

      <q-item-section>
        <q-item-label>{{ currentUser.firstName }} {{ currentUser.lastName }}</q-item-label>
        <q-item-label caption>{{ currentUser.login }}</q-item-label>
      </q-item-section>

      <ProfileFriendshipStatus
          :userId="currentUser.id"
          :friendshipStatus="currentUser.friendshipStatus"
      />
    </q-item>

    <q-separator />

    <q-card-section>
      <ul>
        <li>Age: {{ currentUser.age }}</li>
        <li>City: {{ currentUser.city}}</li>
        <li>Gender: {{ currentUser.gender }}</li>
      </ul>
    </q-card-section>

    <q-separator />

    <q-card-section>
      <div class="text-h6">Interests</div>

      <q-badge
          class="q-ml-sm q-mb-sm"
          rounded color="green"
          v-for="interest in currentUser.interests"
          :label="interest"
      />
    </q-card-section>

<!--    <q-separator />-->

<!--    <q-card-section>Your friend</q-card-section>-->

<!--    <q-card-actions>-->
<!--      <q-btn flat>Add to friend</q-btn>-->
<!--      <q-btn flat color="primary">Reserve</q-btn>-->
<!--    </q-card-actions>-->
  </q-card>
</template>

<style scoped>
</style>
