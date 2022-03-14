import { defineStore } from 'pinia'
import router from "@/router";
import { api } from "@/api";
import type { IApiResponse } from "@/api";
import type { User, IFriendList } from "@/models/user";

function buildAuthHeaders(sessionId: string): Headers {
  const headers = new Headers();

  headers.append('X-SessionId', sessionId);

  return headers;
}

interface IUserStore {
  auth: {
    sessionId: string,
    login: string // TODO: current User
  },
  currentUser: null | User,
  friendList: IFriendList,
  users: User[]
}

export const useUserStore = defineStore({
  // id is required so that Pinia can connect the store to the devtools
  id: 'user',
  state: () =>({
    auth: {
      sessionId: localStorage.getItem("sessionId"),
      login: localStorage.getItem("currLogin"),
    },
    friendList: {
      friends: [],
      requested: [],
      waitingForResponse: [],
    },
    currentUser: null,
    users: [] as User[]
  } as IUserStore),
  getters: {
    sessionId: state => state.auth.sessionId,
    usersList: state => state.users,
  },
  actions:{
    async retrieve(login: string) {
      const user: User = await api.retrieveUser(login, buildAuthHeaders(this.sessionId));

      this.$state.currentUser = user;

      console.log("NEW user", user);
    },

    setSessionId(id: string, login: string) {
      localStorage.setItem("sessionId", id);
      localStorage.setItem("currLogin", login);

      this.$state.auth.login = login;
      this.$state.auth.sessionId = id;
    },

    async login(login: string, password: string) {
      const sessionId = await api.login(login, password); // TODO: OR ERROR

      this.setSessionId(sessionId, login);

      console.log("this.$state.sessionId", this.$state.auth);

      await this.retrieve(login);
      await router.push({name: 'user', params: {login}, replace: true});
    },

    async retrieveUserList() {
      const users: User[] = await api.userList();

      console.log("BEFORE set value: ", users);

      this.$state.users = users;
    },

    async retrieveFriendList(login: string) {
      const friendList: IFriendList = await api.retrieveFriendList(login, buildAuthHeaders(this.sessionId));

      console.log("[friendList]", friendList);

      this.$state.friendList = friendList;
    },

    async userRegister(formData: FormData): Promise<IApiResponse> {
      return api.registerUser(formData)
    },

    async createFriendRequest(targetId: number) {
      console.log("createFriendRequest: targetId: ", targetId);
      // TODO: refresh friendList and userList OR update status by "targetId"
    }
  }
});
