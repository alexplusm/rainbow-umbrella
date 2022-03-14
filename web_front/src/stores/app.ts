import {defineStore} from "pinia";
import {useUserStore} from "@/stores/user";

export const useAppStore = defineStore({
    id: "app",
    state: () => ({}),
    getters: {
        logoNavLink() {
            const userStore = useUserStore();

            if (userStore.currentUser === null) {
                return {name: 'welcome'};
            }
            return {name: 'user', params: {login: userStore.currentUser.login}};
        }
    },
    actions: {}
});
