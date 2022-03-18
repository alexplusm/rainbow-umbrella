<script setup lang="ts">
import ProfileFriendshipStatus from "@/components/ProfileFriendshipStatus.vue";
import type { User } from "@/models/user";
import {useUserStore} from "@/stores/user";

const props = defineProps<{
  user: User;
  showActions?: boolean;
  showApproveFriendshipButton?: boolean;
}>();

const userStore = useUserStore();

const navigationParams = {name: 'user', params: {login: props.user.login}};

function onApproveFriendship(event: PointerEvent) {
  event.preventDefault();
  event.stopPropagation();

  // TODO: or pass friendship ID
  userStore.approveFriendRequest(props.user.id);
  console.log("[onApproveFriendship]");
}
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

    <template v-if="props.showApproveFriendshipButton">
      <q-separator vertical />

      <q-item-section side>
        <q-btn flat rounded icon="done" class="text-green" v-on:click="onApproveFriendship"/>
      </q-item-section>
    </template>
  </q-item>
</template>

<style scoped>
</style>
