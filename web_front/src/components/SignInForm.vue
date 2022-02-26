<script>
import { ref } from 'vue'
import { useQuasar } from 'quasar'
import router from '@/router'

export default {
  setup() {
    const $q = useQuasar()

    const login = ref(null)
    const password = ref(null)

    return {
      login,
      password,

      onSubmit () {
        const formData = {
          login: login.value,
          password: password.value
        };

        fetch("/api/v1/users/login", {method: "POST", body: JSON.stringify(formData)})
          .then(response => response.json())
          .then(data => {
              const sessionId = data["sessionID"]

              console.log("sessionId: ", sessionId);
              localStorage.setItem("X-SessionId", sessionId);
              localStorage.setItem("currUser", login.value);
            });

        router.push({name: 'home', params: {login: login.value}, replace: true});
      }
    }
  }
}
</script>

<template>
  <div class="wrap q-pa-md" style="max-width: 400px">

    <q-card class="my-card">
      <q-card-section>
        <div class="text-h6">Sign In</div>
      </q-card-section>

      <q-separator dark />

      <q-card-actions>
        <q-form @submit="onSubmit" class="q-gutter-md">
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

          <div>
            <q-btn label="Submit" type="submit" color="primary"/>
          </div>
        </q-form>
      </q-card-actions>
    </q-card>
  </div>
</template>

<style scoped>
.wrap {
  display: flex;
  flex-direction: column;
}
</style>