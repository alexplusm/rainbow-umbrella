<script>
import { useQuasar } from 'quasar'
import { ref } from 'vue'

export default {
  setup () {
    const $q = useQuasar()

    const form = ref(null)

    // function validate () {
    //   form.value.validate().then(success => {
    //     if (success) {
    //       // yay, models are correct
    //       console.log("SUCCESS");
    //     }
    //     else {
    //       console.log("FAILURE");
    //     }
    //   })
    // }

    const firstName = ref(null)
    const lastName = ref(null)
    const birthday = ref(null)

    const age = ref(null)
    const accept = ref(false)

    console.log("$q", $q);

    return {
      form,

      firstName,
      lastName,
      birthday,

      age,
      accept,

      onSubmit () {
        // console.log("form.value", form.value.validate());
        // console.log("form.value", form.value);

        $q.notify({
          color: 'green-4',
          textColor: 'white',
          icon: 'cloud_done',
          message: 'Register succeed'
        })
      },

      onReset () {
        firstName.value = null
        age.value = null
        accept.value = false
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
          v-model="firstName"
          label="First name"
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

<!--      <q-input-->
<!--          filled-->
<!--          type="number"-->
<!--          v-model="age"-->
<!--          label="Your age *"-->
<!--          lazy-rules-->
<!--          :rules="[-->
<!--          val => val !== null && val !== '' || 'Please type your age',-->
<!--          val => val > 0 && val < 100 || 'Please type a real age'-->
<!--        ]"-->
<!--      />-->

      <div>
        <q-btn label="Submit" type="submit" color="primary"/>
        <q-btn label="Reset" type="reset" color="primary" flat class="q-ml-sm" />
      </div>
    </q-form>
  </div>
</template>

<style scoped>
</style>
