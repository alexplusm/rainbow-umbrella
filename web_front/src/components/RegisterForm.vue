<script setup lang="ts">
import { ref } from 'vue'
import { useQuasar } from 'quasar'
import { useUserStore } from "@/stores/user";

const $q = useQuasar();
const userStore = useUserStore();

const login = ref("login1")
const password = ref("")
const firstName = ref("guest")
const lastName = ref("guestov")
const birthday = ref("1996/10/17")
const gender = ref("Male")
const city = ref("Moscow")
const interests = ref([])

function onReset() {
  login.value = "";
  password.value = "";
  firstName.value = "";
  lastName.value = "";
  birthday.value = "";
  gender.value = "";
  city.value = "";
  interests.value = [];
}

function onSubmit () {
  const formData = new FormData()
  formData.append("login", login.value)
  formData.append("password", password.value)
  formData.append("firstName", firstName.value)
  formData.append("lastName", lastName.value)
  formData.append("birthday", birthday.value)
  formData.append("gender", gender.value)
  formData.append("city", city.value)
  formData.append("interests", interests.value.join(","))

  userStore.userRegister(formData).then(apiResp => {
    if (!apiResp.hasError) {
      $q.notify({
        color: 'green-4',
        textColor: 'white',
        icon: 'cloud_done',
        message: 'Register succeed'
      });
      onReset();
    } else {
      $q.notify({
        color: 'red-4',
        textColor: 'white',
        icon: 'cloud_done',
        message: apiResp.notifyMessage as string
      });
    }
  });
}

function addInterest (val: string, done: any) {
  done(val, 'add-unique');
}
</script>

<template>
  <div>
    <q-form
        @submit="onSubmit"
        @reset="onReset"
        class="q-gutter-md"
    >
      <div class="row">
        <q-input
            class="col"
            filled
            v-model="login"
            label="Login"
            lazy-rules
            :rules="[ val => val && val.length > 0 || 'Please type something']"
        />
        <q-input
            filled
            class="col"
            v-model="password"
            label="Password"
            type="password"
            lazy-rules
            :rules="[ val => val && val.length > 0 || 'Please type something']"
        />
      </div>

      <div class="row">
        <q-input
            filled
            class="col"
            v-model="firstName"
            label="First name"
            lazy-rules
            :rules="[ val => val && val.length > 0 || 'Please type something']"
        />
        <q-input
            filled
            class="col"
            v-model="lastName"
            label="Last name"
            lazy-rules
            :rules="[ val => val && val.length > 0 || 'Please type something']"
        />
      </div>

      <q-input filled label="Birthday" v-model="birthday" mask="date" :rules="['date']">
        <template v-slot:append>
          <q-icon name="event" class="cursor-pointer">
            <q-popup-proxy ref="qDateProxy" cover transition-show="scale" transition-hide="scale">
              <q-date label="Birthday" v-model="birthday">
                <div class="row items-center justify-end">
                  <q-btn v-close-popup label="Close" color="primary" flat />
                </div>
              </q-date>
            </q-popup-proxy>
          </q-icon>
        </template>
      </q-input>

      <q-select
          filled
          v-model="gender"
          label="Gender"
          :options="['Male', 'Female']"
          lazy-rules
          :rules="[ val => val && val.length > 0 || 'Please select something']"
      />
      <q-input
          filled
          v-model="city"
          label="City"
          lazy-rules
          :rules="[ val => val && val.length > 0 || 'Please type something']"
      />
      <q-select
          label="Interests"
          filled
          v-model="interests"
          use-input
          use-chips
          multiple
          hide-dropdown-icon
          input-debounce="0"
          @new-value="addInterest"
      />

      <div>
        <q-btn label="Submit" type="submit" color="primary"/>
        <q-btn label="Reset" type="reset" color="primary" flat class="q-ml-sm" />
      </div>
    </q-form>
  </div>
</template>

<style scoped>

.col + .col {
  margin-left: 1rem;
}
</style>
