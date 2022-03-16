import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from "@/stores/user";
import WelcomeView from '@/views/WelcomeView.vue'
import UserView from '@/views/UserView.vue'
import NotFoundView from '@/views/NotFoundView.vue'
import {useAuthStore} from "@/stores/auth";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'welcome',
      component: WelcomeView
    },
    {
      path: '/not-found',
      name: 'notFound',
      component: NotFoundView
    },
    {
      path: '/:login',
      name: 'user',
      component: UserView
    },
  ]
});

router.beforeResolve(async (to, from, next) => {
  const userStore = useUserStore();
  const authStore = useAuthStore();

  await authStore.setUser();

  console.log("router.beforeResolve: route name: ", to.name);

  if (to.name === 'user') {
    const login = to.params['login'] as string;
    console.log("router.beforeResolve: user route: ", login);

    await userStore.uploadDataForUser(login);
  }

  next();
})

export default router
