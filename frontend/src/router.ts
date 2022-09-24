import { createRouter, createWebHashHistory } from 'vue-router'
import Home from './components/Home.vue'
import { useCounterStore } from './store'

const routes = [
  {
    path: '/home',
    name: 'Home',
    component: Home,
  },
  {
    path: '/features',
    name: 'Features',
    component: () => import('./components/Features.vue'),
  },
  {
    path: '/',
    name: 'Start',
    component: () => import('./components/Start.vue'),
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

// router.beforeEach((to) => {
//   const counterStore = useCounterStore()
// })

export default router
