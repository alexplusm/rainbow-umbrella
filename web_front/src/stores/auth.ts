import {defineStore} from "pinia";
import router from "@/router";
import { api } from "@/api";
import type { IApiResponse } from "@/api";

interface IAuthStore {
    sessionId: string;
    login: string;
}

export const useAuthStore = defineStore({
    id: 'auth',
    state: () => ({
        sessionId: localStorage.getItem("sessionId"),
        login: localStorage.getItem("login")
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

            await router.push({name: 'user', params: {login}, replace: true});

            return response;
        },

        async logout() {
            localStorage.removeItem("sessionId");
            localStorage.removeItem("login");

            await router.push({name: 'welcome', replace: true});

            api.logout()
        },

        setSessionId(id: string, login: string) {
            localStorage.setItem("sessionId", id);
            localStorage.setItem("login", login);

            this.$state.login = login;
            this.$state.sessionId = id;
        }
    }
})