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
</script>

<template>
  <q-item clickable v-ripple>

    <q-item-section avatar>
      <q-avatar>
        <img src="https://cdn.quasar.dev/img/boy-avatar.png">
      </q-avatar>
    </q-item-section>

    <q-item-section>{{ props.user.firstName }} {{ props.user.lastName }}</q-item-section>



    <q-separator vertical></q-separator>

    ({{ props.showActions }})

    <q-item-section side>
<!--      <q-item-label caption>5 min ago</q-item-label>-->
<!--      <q-icon size="2em" name="star" color="yellow" />-->

      <q-btn v-if="showAddButton()" flat round icon="add" class="text-blue" />
      <q-icon v-if="showFriendIcon()" size="3em" class="text-green" name="grade"></q-icon>
      <q-icon v-if="showWaitIcon()" size="3em" class="text-yellow" name="watch_later"></q-icon>


    </q-item-section>

<!--    <q-item-section>-->
<!--      <q-icon color="primary" name="bluetooth" />-->
<!--    </q-item-section>-->
  </q-item>
</template>


<style scoped>
</style>
