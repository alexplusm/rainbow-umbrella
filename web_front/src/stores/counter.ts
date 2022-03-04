import { defineStore } from 'pinia'
import {UserM} from "@/models/user";
import router from "@/router";

export const useCounterStore = defineStore({
  id: 'counter',
  state: () => ({
    counter: 0
  }),
  getters: {
    doubleCount: (state) => state.counter * 2
  },
  actions: {
    increment() {
      this.counter++
    }
  }
})



// import { defineStore } from 'pinia'

export const useUserStore = defineStore({
  // id is required so that Pinia can connect the store to the devtools
  id: 'user',
  state: () =>({
    sessionId: null
  }),
  getters: {},
  actions:{
    // async retrieve(login: string) {
    //   const user: UserM = await fetch(`/api/v1/users/${login}`, {headers})
    //       .then(resp => resp.json())
    //       .then(data => {
    //         console.log("data", data);
    //         const user = new UserM(data)
    //
    //         user.age = 666;
    //
    //         return user;
    //       });
    // }

    async login(login: string, password: string) {
      const data = {login, password}

      await fetch("/api/v1/users/login", {method: "POST", body: JSON.stringify(data)})
        .then(response => response.json())
        .then(data => {
          const sessionId = data["sessionID"]

          console.log("sessionId: ", sessionId);
          localStorage.setItem("X-SessionId", sessionId);
          localStorage.setItem("currUser", login);

          this.$state.sessionId = sessionId;
        });

      console.log("this.$state.sessionId", this.$state.sessionId);


      await router.push({name: 'home', params: {login}, replace: true});
    }
  }
})