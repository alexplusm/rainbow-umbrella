<script>
import { useQuasar } from 'quasar'
import { ref } from 'vue'

export default {
  setup () {
    const $q = useQuasar()

    const form = ref(null)

    const login = ref("login1")
    const password = ref("pass")

    const firstName = ref("alex")
    const lastName = ref("kekus")
    const birthday = ref("1996/10/17")
    const gender = ref("Male")
    const city = ref("Moscow")
    // TODO: interests

    return {
      form,

      login,
      password,

      firstName,
      lastName,
      birthday,
      gender,
      city,

      onSubmit () {
        // console.log("form.value", form.value.validate());
        // console.log("form.value", form.value);

        const formData = new FormData()
        formData.append("login", login.value)
        formData.append("password", password.value)
        formData.append("firstName", firstName.value)
        formData.append("lastName", lastName.value)
        formData.append("birthday", birthday.value)
        formData.append("gender", gender.value)
        formData.append("city", city.value)

        fetch("/api/v1/user/register", {
          method: "POST",
          body: formData
        }).then(response => {
          switch (response.status) {
            case 201:
              $q.notify({
                color: 'green-4',
                textColor: 'white',
                icon: 'cloud_done',
                message: 'Register succeed'
              });
              break;
            case 409:
              $q.notify({
                color: 'red-4',
                textColor: 'white',
                icon: 'cloud_done',
                message: 'Login already taken'
              });
              break;
            default:
              $q.notify({
                color: 'red-4',
                textColor: 'white',
                icon: 'cloud_done',
                message: 'Somethings went wrong'
              });
          }
        })

        // $q.notify({
        //   color: 'green-4',
        //   textColor: 'white',
        //   icon: 'cloud_done',
        //   message: 'Register succeed'
        // })
        // TODO: send request
        //    process 409 - unique login
      },

      onReset () {
        login.value = null
        password.value = null

        firstName.value = null
        lastName.value = null
        birthday.value = null
        gender.value = null
        city.value = null
      }
    }
  }
}
</script>

<template>
  <div class="q-pa-md" style="max-width: 400px">
    <q-form
        ref="form"
        @submit="onSubmit"
        @reset="onReset"
        class="q-gutter-md"
    >
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

      <q-input
          filled
          v-model="firstName"
          label="First name"
          lazy-rules
          :rules="[ val => val && val.length > 0 || 'Please type something']"
      />
      <q-input
          filled
          v-model="lastName"
          label="Last name"
          lazy-rules
          :rules="[ val => val && val.length > 0 || 'Please type something']"
      />
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

      <div>
        <q-btn label="Submit" type="submit" color="primary"/>
        <q-btn label="Reset" type="reset" color="primary" flat class="q-ml-sm" />
      </div>
    </q-form>
  </div>
</template>

<style scoped>
</style>
