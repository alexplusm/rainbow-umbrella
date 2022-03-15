import {defineStore} from "pinia";
import type {User} from "@/models/user";
import { api } from "@/api";
import type { IApiResponse } from "@/api";
import router from "@/router";


interface IAuthStore {
    sessionId: string;
    login: string;
    user: null | User;
}

export const useAuthStore = defineStore({
    id: 'auth',
    state: () => ({
        sessionId: localStorage.getItem("sessionId"),
        login: localStorage.getItem("login"),
        user: null,
    } as IAuthStore),
    getters: {},
    actions: {
        async login(login: string, password: string): Promise<IApiResponse<string>> {
            const response = await api.login(login, password); // TODO: OR ERROR

            if (response.hasError) {
                return response;
            }

            this.setSessionId(response.data, login);

            // TODO: !!!
            // await this.retrieve(login);
            // await router.push({name: 'user', params: {login}, replace: true});

            return response;
        },

        setSessionId(id: string, login: string) {
            localStorage.setItem("sessionId", id);
            localStorage.setItem("login", login);

            this.$state.login = login;
            this.$state.sessionId = id;
        },

    }
})