import { defineStore } from 'pinia'
import {UserM} from "@/models/user";
import router from "@/router";
import {loginApi} from "@/api";

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

function buildHeaders(sessionId: string): Headers {
  const headers = new Headers();

  headers.append('X-SessionId', sessionId)

  return headers;
}

export const useUserStore = defineStore({
  // id is required so that Pinia can connect the store to the devtools
  id: 'user',
  state: () =>({
    auth: {
      sessionId: "",
      login: "",
    }
  }),
  getters: {
    sessionId: state => state.auth.sessionId
  },
  actions:{
    async retrieve(login: string) {
      const headers = buildHeaders(this.sessionId);

      const user: UserM = await fetch(`/api/v1/users/${login}`, {headers})
          .then(resp => resp.json())
          .then(data => {
            console.log("data", data);
            const user = new UserM(data)

            user.age = 666;

            return user;
          });

      console.log("NEW USEEERRR", user);
    },

    setSessionId(id: string, login: string) {
      localStorage.setItem("sessionId", id);
      localStorage.setItem("currLogin", login);

      this.$state.auth.login = login;
      this.$state.auth.sessionId = id;
    },

    async login(login: string, password: string) {
      const sessionId = await loginApi(login, password);

      this.setSessionId(sessionId, login);
      console.log("this.$state.sessionId", this.$state.auth);

      await router.push({name: 'home', params: {login}, replace: true});
    }
  }
})