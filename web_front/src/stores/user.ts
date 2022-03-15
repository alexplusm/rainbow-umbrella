import { defineStore } from 'pinia'
import { api } from "@/api";
import type { IApiResponse } from "@/api";
import type { User, IFriendList } from "@/models/user";
import { useAuthStore } from "@/stores/auth";


interface IUserStore {
  auth: {
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
    usersList: state => state.users
  },
  actions:{
    async retrieve(login: string) {
      const user = await api.retrieveUser(login, useAuthStore().authHeaders);

      // TODO: RM!
      this.$state.currentUser = user;

      console.log("NEW user", user);
    },

    async retrieveUserList() {
      const users: User[] = await api.userList();

      console.log("BEFORE set value: ", users);

      this.$state.users = users;
    },

    async retrieveFriendList(login: string) {
      const friendList = await api.retrieveFriendList(login, useAuthStore().authHeaders);

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
