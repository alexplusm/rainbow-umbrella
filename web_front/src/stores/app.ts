import { defineStore } from "pinia";
import { useAuthStore } from "@/stores/auth";

export const useAppStore = defineStore({
    id: "app",
    state: () => ({}),
    getters: {
        logoNavLink() {
            const authStore = useAuthStore()

            if (authStore.user === null) {
                return {name: 'welcome'};
            }
            return {name: 'user', params: {login: authStore.login}};
        }
    },
    actions: {}
});
