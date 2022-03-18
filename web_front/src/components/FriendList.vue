<script setup lang="ts">
import { ref } from "vue";
import { useUserStore } from "@/stores/user";
import ProfileMini from "@/components/ProfileMini.vue";

const userStore = useUserStore();
const tab = ref('friends');
</script>

<template>
  <section class="wrap">
    <div class="text-h6 q-pb-sm">Friend List</div>

    <q-tabs
        v-model="tab"
        dense
        class="text-grey"
        active-color="primary"
        indicator-color="primary"
        align="justify"
        narrow-indicator
    >
      <q-tab name="friends" label="Friends">
        <q-badge color="green" floating>{{ userStore.$state.friendList.friends.length }}</q-badge>
      </q-tab>
      <q-tab name="requested" label="Requested">
        <q-badge color="blue" floating>{{ userStore.$state.friendList.requested.length }}</q-badge>
      </q-tab>
      <q-tab name="waiting" label="Waiting">
        <q-badge color="amber" floating>{{ userStore.$state.friendList.waitingForResponse.length }}</q-badge>
      </q-tab>
    </q-tabs>

    <q-separator />

    <q-tab-panels v-model="tab" animated>
      <q-tab-panel name="friends">
        <div v-if="userStore.$state.friendList.friends.length === 0">
          Empty
        </div>

        <q-list
          v-if="userStore.$state.friendList.friends.length !== 0"
          class="user_list"
          bordered
          separator
        >
          <ProfileMini
            v-for="user in userStore.$state.friendList.friends"
            :key="user.id"
            class="profile_mini"
            :user="user"
            :showActions="false"
          />
        </q-list>
      </q-tab-panel>

      <q-tab-panel name="requested">
        <div v-if="userStore.$state.friendList.requested.length === 0">
          Empty
        </div>

        <q-list
          v-if="userStore.$state.friendList.requested.length !== 0"
          class="user_list"
          bordered
          separator
        >
          <ProfileMini
            v-for="user in userStore.$state.friendList.requested"
            :key="user.id"
            class="profile_mini"
            :user="user"
            :showActions="false"
          />
        </q-list>
      </q-tab-panel>

      <q-tab-panel name="waiting">
        <div v-if="userStore.$state.friendList.waitingForResponse.length === 0">
          Empty
        </div>

        <q-list
          v-if="userStore.$state.friendList.waitingForResponse.length !== 0"
          class="user_list"
          bordered
          separator
        >
          <ProfileMini
            v-for="item in userStore.$state.friendList.waitingForResponse"
            class="profile_mini"
            :user="item"
            :showActions="false"
          />
        </q-list>
      </q-tab-panel>
    </q-tab-panels>
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
