import { defineStore } from 'pinia'
import router from "@/router";
import { api } from "@/api";
import type { IApiResponse } from "@/api";
import type { User, IFriendList } from "@/models/user";
import { useAuthStore } from "@/stores/auth";
import { FriendshipStatus } from "@/models/user";


interface IUserStore {
  currentUser: User;
  friendList: IFriendList;
  users: User[];
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
      const userResponse = await this.retrieve(login);
      if (userResponse.hasError) {
        await useAuthStore().logout()
        return router.push({name: 'welcome'});
      }
      this.$state.currentUser = userResponse.data;

      await this.retrieveFriendList(login);
      await this.retrieveUserList();
    },

    async retrieve(login: string): Promise<IApiResponse<User>> {
      return await api.retrieveUser(login, useAuthStore().authHeaders);
    },

    async retrieveUserList() {
      this.$state.users = await api.userList(useAuthStore().authHeaders);
    },

    async retrieveFriendList(login: string) {
      this.$state.friendList = await api.retrieveFriendList(login, useAuthStore().authHeaders);
    },

    async userRegister(formData: FormData): Promise<IApiResponse> {
      return api.registerUser(formData)
    },

    async createFriendRequest(targetId: number) {
      const response = await api.createFriendship(useAuthStore().user.id, targetId);

      if (response.hasError) {
        console.log("[createFriendRequest]: error: ", response);
      } else {
        if (this.$state.currentUser.id !== useAuthStore().user.id) {
          return
        }

        this.$state.users = this.$state.users.map(user => {
          if (user.id === targetId) {
            user.friendshipStatus = FriendshipStatus.Wait;
            this.$state.friendList.waitingForResponse.push(user);
          }

          return user;
        });
      }
    },

    async approveFriendRequest(targetId: number) {
      // TODO: move from waiting to friends
    },

    async declineFriendRequest(targetId: number) {
      // TODO: rm from waiting
    }
  }
});
