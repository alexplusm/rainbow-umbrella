import { defineStore } from "pinia";
import { useAuthStore } from "@/stores/auth";

export const useAppStore = defineStore({
    id: 'app',
    state: () => ({}),
    getters: {
        logoNavLink() {
            const authStore = useAuthStore()

            if (authStore.$state.login === "") {
                return {name: 'welcome'};
            }
            return {name: 'user', params: {login: authStore.$state.login}};
        }
    },
    actions: {}
});
