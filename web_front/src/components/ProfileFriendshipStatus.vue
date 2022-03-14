<script setup lang="ts">
import { FriendshipStatus } from "@/models/user";

const props = defineProps<{userId: number, friendshipStatus: FriendshipStatus}>();

function showAddButton(): boolean {
  return props.friendshipStatus === FriendshipStatus.NotFriend;
}

function showWaitIcon(): boolean {
  return props.friendshipStatus === FriendshipStatus.Wait;
}

function showFriendIcon(): boolean {
  return props.friendshipStatus === FriendshipStatus.Friend;
}

function showMyselfIcon(): boolean {
  return props.friendshipStatus === FriendshipStatus.Myself;
}

function onAddFriend(event: PointerEvent) {
  event.stopPropagation();
  event.preventDefault();

  // TODO: store.initFriendship();

  console.log("onAddFriend: ", props.userId);
}

</script>

<template>
  <q-item-section side>
    <q-btn v-on:click="onAddFriend" v-if="showAddButton()" flat round icon="add" class="text-blue" />

    <q-icon v-if="showFriendIcon()" size="3em" class="text-green" name="grade" />

    <q-icon v-if="showWaitIcon()" size="3em" class="text-yellow" name="watch_later" />

    <q-icon v-if="showMyselfIcon()" size="3em" class="text-blue" name="self_improvement" />
  </q-item-section>
</template>

<style scoped>
</style>
