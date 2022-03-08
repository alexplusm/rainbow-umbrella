import { createRouter, createWebHistory } from 'vue-router'
import WelcomeView from '@/views/WelcomeView.vue'
import UserView from '@/views/UserView.vue'
import NotFoundView from '@/views/NotFoundView.vue'
import {useUserStore} from "@/stores/user";

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

router.beforeResolve((to, from, next) => {
  const userStore = useUserStore();

  console.log("router.beforeResolve: route name: ", to.name);

  if (to.name === 'user') {
    const login = to.params['login'] as string;

    console.log("router.beforeResolve: user route: ", login);

    userStore.retrieve(login)
        .then(() => userStore.retrieveFriendList(login))
        .then(() => next());
  } else {
    next();
  }
})

export default router
