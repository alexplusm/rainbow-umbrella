import { defineStore } from 'pinia'
import { api } from "@/api";
import type { IApiResponse } from "@/api";
import type { User, IFriendList } from "@/models/user";
import { useAuthStore } from "@/stores/auth";


interface IUserStore {
  currentUser: User,
  friendList: IFriendList,
  users: User[]
}

export const useUserStore = defineStore({
  id: 'user',
  state: () =>({
    currentUser: {} as User,
    friendList: {
      friends: [],
      requested: [],
      waitingForResponse: [],
    },
    users: [] as User[]
  } as IUserStore),
  getters: {
    usersList: state => state.users
  },
  actions:{
    async uploadDataForUser(login: string) {
      this.$state.currentUser = await this.retrieve(login);
      await this.retrieveFriendList(login);
      await this.retrieveUserList();
    },

    async retrieve(login: string): Promise<User> {
      return await api.retrieveUser(login, useAuthStore().authHeaders);
    },

    async retrieveUserList() {
      this.$state.users = await api.userList();
    },

    async retrieveFriendList(login: string) {
      this.$state.friendList = await api.retrieveFriendList(login, useAuthStore().authHeaders);
    },

    async userRegister(formData: FormData): Promise<IApiResponse> {
      return api.registerUser(formData)
    },

    async createFriendRequest(targetId: number) {
      useAuthStore()

      const response = await api.createFriendship(0, targetId);

      if (response.hasError) {
        console.log("[createFriendRequest]: error: ", response);
      } else {
        // TODO: refresh friendList and userList OR update status by "targetId"
      }
    }
  }
});
