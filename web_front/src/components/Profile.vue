<template>
  <section>
    KEKW
<!--    {{ user.login }}-->
  </section>
</template>

<script lang="ts" setup>
import router from '@/router'
import {UserM} from "@/models/user";
import {onUnmounted} from "vue";


const login = router.currentRoute.value.params['login'];
console.log("curr login", login);

const sessionId = localStorage.getItem("X-SessionId");

console.log("sessionId", sessionId)

// withAsyncContext()

const headers = new Headers();
if (sessionId != null) {
  headers.append('X-SessionId', sessionId)
}

const user: UserM = await fetch(`/api/v1/users/${login}`, {headers})
  .then(resp => resp.json())
  .then(data => {
    console.log("data", data);
    const user = new UserM(data)

    user.age = 666;

    return user;
  });

console.log("USERRR: ", user);


// export default {
//   async setup() {
//     const login = router.currentRoute.value.params['login'];
//     console.log("curr login", login);
//
//     const sessionId = localStorage.getItem("X-SessionId");
//
//     console.log("sessionId", sessionId)
//
//     // withAsyncContext()
//
//     const headers = new Headers();
//     if (sessionId != null) {
//       headers.append('X-SessionId', sessionId)
//     }
//
//     // const user: UserM = await fetch(`/api/v1/users/${login}`, {headers})
//     //   .then(resp => resp.json())
//     //   .then(data => {
//     //     console.log("data", data);
//     //     const user = new UserM(data)
//     //
//     //     user.age = 666;
//     //
//     //     return user;
//     //   });
//     //
//     // console.log("USERRR: ", user);
//
//     onUnmounted(
//         () => console.log('Unmounted'),
//         // instance // <--- pass the instance to it
//         // "wtf";
//     )
//
//     return {
//       // user
//     }
//   }
// }
</script>

<style scoped>
</style>