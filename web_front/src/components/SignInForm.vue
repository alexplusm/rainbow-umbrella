<script setup lang="ts">
import { ref } from 'vue'
import { useQuasar } from 'quasar'
import { useUserStore } from "@/stores/user";

const $q = useQuasar();
const userStore = useUserStore();

const login = ref("");
const password = ref("");

function onSubmit () {
  userStore.login(login.value, password.value).then(data => {
    console.log("data: ", data);
    // TODO: process ERROR
  });
}
</script>

<template>
  <q-card style="height: 280px">
    <q-card-section>
      <div class="text-h6 q-mb-lg">Sign In</div>

      <q-form @submit="onSubmit">
        <q-input
          filled
          v-model="login"
          label="Login"
          lazy-rules
          :rules="[ val => val && val.length > 0 || 'Please type something']"
        />

        <q-input
          filled
          v-model="password"
          label="Password"
          type="password"
          lazy-rules
          :rules="[ val => val && val.length > 0 || 'Please type something']"
        />

        <q-btn label="Submit" type="submit" color="primary"/>
      </q-form>
    </q-card-section>
  </q-card>
</template>

<style scoped>
</style>
