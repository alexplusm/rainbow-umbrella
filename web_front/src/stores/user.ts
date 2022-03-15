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
    sessionId: string;
    login: string; // TODO: current User
    currentUser: User;
  },
  currentUser: User,
  chosenUser: User,
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
      currentUser: {} as User,
    },
    friendList: {
      friends: [],
      requested: [],
      waitingForResponse: [],
    },
    currentUser: {} as User, // TODO: remove
    chosenUser: {} as User,
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

    // setSessionId(id: string, login: string) {
    //   localStorage.setItem("sessionId", id);
    //   localStorage.setItem("currLogin", login);
    //
    //   this.$state.auth.login = login;
    //   this.$state.auth.sessionId = id;
    // },

    // async login(login: string, password: string): Promise<IApiResponse<string>> {
    //   const response = await api.login(login, password); // TODO: OR ERROR
    //
    //   if (response.hasError) {
    //     return response;
    //   }
    //
    //   this.setSessionId(response.data, login);
    //   await this.retrieve(login);
    //   await router.push({name: 'user', params: {login}, replace: true});
    //
    //   return response;
    // },

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
      const response = await api.createFriendship(0, targetId);

      if (response.hasError) {
        console.log("[createFriendRequest]: error: ", response);
      } else {
        // TODO: refresh friendList and userList OR update status by "targetId"
      }
    }
  }
});
