<script setup lang="ts">
import type {User} from "@/models/user";

const props = defineProps<{user: User, showActions: boolean}>();

function showAddButton(): boolean {
  return props.user.friendshipStatus === "new";
}

function showWaitIcon(): boolean {
  return props.user.friendshipStatus === "wait";
}

function showFriendIcon(): boolean {
  return props.user.friendshipStatus === "friend";
}
function onAddFriend(event: PointerEvent) {
  console.log("event: ", event);
  event.stopPropagation();
}
</script>

<template>
  <q-item clickable v-ripple :to="{name: 'user', params: {login: props.user.login}}">
    <q-item-section avatar>
      <q-avatar >
        <img alt="avatar" src="https://cdn.quasar.dev/img/boy-avatar.png">
      </q-avatar>
    </q-item-section>

    <q-item-section>{{ props.user.firstName }} {{ props.user.lastName }}</q-item-section>

    <template v-if="props.showActions">
      <q-separator vertical></q-separator>

      <q-item-section side>
        <q-btn v-on:click="onAddFriend" v-if="showAddButton()" flat round icon="add" class="text-blue" />
        <q-icon v-if="showFriendIcon()" size="3em" class="text-green" name="grade"></q-icon>
        <q-icon v-if="showWaitIcon()" size="3em" class="text-yellow" name="watch_later"></q-icon>
      </q-item-section>
    </template>
  </q-item>
</template>


<style scoped>
</style>
