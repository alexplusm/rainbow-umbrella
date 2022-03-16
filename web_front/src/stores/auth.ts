import { defineStore } from "pinia";
import router from "@/router";
import { api } from "@/api";
import type { IApiResponse } from "@/api";
import type {User} from "@/models/user";
import { useUserStore } from "@/stores/user";

interface IAuthStore {
    sessionId: string;
    login: string;
    user: User;
}

export const useAuthStore = defineStore({
    id: 'auth',
    state: () => ({
        sessionId: localStorage.getItem("sessionId"),
        login: localStorage.getItem("login"),
        user: {}
    } as IAuthStore),
    getters: {
        authHeaders: (state): Headers => new Headers({'X-SessionId': state.sessionId})
    },
    actions: {
        async login(login: string, password: string): Promise<IApiResponse<string>> {
            const response = await api.login(login, password); // TODO: OR ERROR

            if (response.hasError) {
                return response;
            }

            this.setSessionId(response.data, login);
            await this.setUser();

            await router.push({name: 'user', params: {login}, replace: true});

            return response;
        },

        async logout() {
            localStorage.removeItem("sessionId");
            localStorage.removeItem("login");

            await api.logout();
            await router.push({name: 'welcome', replace: true});
        },

        async setUser() {
            if (!this.$state.user.id) {
                this.$state.user = await useUserStore().retrieve(this.$state.login);
            }
        },

        setSessionId(id: string, login: string) {
            localStorage.setItem("sessionId", id);
            localStorage.setItem("login", login);

            this.$state.login = login;
            this.$state.sessionId = id;
        }
    }
})