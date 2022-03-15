<script setup lang="ts">
import ProfileFriendshipStatus from "@/components/ProfileFriendshipStatus.vue";
import type { User } from "@/models/user";

const props = defineProps<{user: User, showActions: boolean}>();
const navigationParams = {name: 'user', params: {login: props.user.login}};
</script>

<template>
  <q-item clickable v-ripple :to="navigationParams">
    <q-item-section avatar>
      <q-avatar>
        <img alt="avatar" :src="props.user.avatarUrl">
      </q-avatar>
    </q-item-section>

    <q-item-section>{{ props.user.firstName }} {{ props.user.lastName }}</q-item-section>

    <template v-if="props.showActions">
      <q-separator vertical />

      <ProfileFriendshipStatus
        :userId="props.user.id"
        :friendshipStatus="props.user.friendshipStatus"
      />
    </template>
  </q-item>
</template>

<style scoped>
</style>
