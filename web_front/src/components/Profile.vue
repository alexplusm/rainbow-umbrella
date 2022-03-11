<script lang="ts" setup>
import { useUserStore } from "@/stores/user";
import type { User } from "@/models/user";

const userStore = useUserStore();
const currentUser: User = userStore.currentUser as User;
</script>

<template>
  <section>
    <q-card flat bordered>
      <q-item>
        <q-item-section avatar>
          <q-avatar>
            <img src="https://cdn.quasar.dev/img/boy-avatar.png">
          </q-avatar>
        </q-item-section>

        <q-item-section>
          <q-item-label>{{ currentUser.firstName }} {{ currentUser.lastName }}</q-item-label>
          <q-item-label caption>
            {{ currentUser.login }}
          </q-item-label>
        </q-item-section>
      </q-item>

      <q-separator />

      <q-card-section>
        <ul>
          <li>
            Age: {{ currentUser.age }}
          </li>
          <li>
            City: {{ currentUser.city}}
          </li>
          <li>
            Gender: {{ currentUser.gender }}
          </li>
        </ul>
      </q-card-section>

      <q-separator />

      <q-card-section>
        Your friend
      </q-card-section>

      <q-card-actions>
        <q-btn flat>
          Add to friend
        </q-btn>
        <q-btn flat color="primary">
          Reserve
        </q-btn>
      </q-card-actions>
    </q-card>
  </section>
</template>

<style scoped>
</style>
