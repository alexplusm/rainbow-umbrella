import { defineStore } from "pinia";
import { useAuthStore } from "@/stores/auth";

export const useAppStore = defineStore({
    id: 'app',
    state: () => ({}),
    getters: {
        logoNavLink() {
            const login = useAuthStore().$state.login

            if (login === "" || login === null) {
                return {name: 'welcome'};
            }
            return {name: 'user', params: {login}};
        }
    },
    actions: {}
});
