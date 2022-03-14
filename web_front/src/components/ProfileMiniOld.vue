<script setup lang="ts">
import type { User } from "@/models/user";

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
  <div>
    <q-card class="card" flat bordered>
      <q-card-section horizontal>
        <template v-if="props.showActions">
          <q-card-actions class="actions">
            <q-btn v-if="showAddButton()" flat round icon="add" class="text-blue" />
            <q-icon v-if="showFriendIcon()" size="3em" class="text-green" name="grade"></q-icon>
            <q-icon v-if="showWaitIcon()" size="3em" class="text-yellow" name="watch_later"></q-icon>
          </q-card-actions>

          <q-separator vertical />
        </template>

        <q-item>
          <q-item-section avatar>
            <q-avatar>
              <img src="https://cdn.quasar.dev/img/boy-avatar.png">
            </q-avatar>
          </q-item-section>

          <q-item-section>
            <q-item-label class="user_name ellipsis">
              {{ props.user.firstName }} {{ props.user.lastName }}
            </q-item-label>
          </q-item-section>
        </q-item>
      </q-card-section>
    </q-card>
  </div>
</template>

<style scoped>
.user_name {
  width: 200px;
}
</style>
