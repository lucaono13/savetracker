import { createRouter, createWebHashHistory } from 'vue-router'
import Home from './components/Pages/Home.vue'
import { useCounterStore } from './stores/counter'
import { useSavesStore } from './stores/saves'
import { PrimeIcons } from 'primevue/api'

// import TrackerFeatures from './components/TrackerFeatures.vue'

const routes = [
  {
    path: '/home',
    name: 'Home 2',
    component: Home,
    meta: {
      icon: "fa-solid fa-mug-hot",
      secondary: false
    },
    
  },
  {
    path: '/features',
    name: 'Features',
    component: () => import('./components/Pages/Features.vue'),
    meta: {
      icon: PrimeIcons.DISCORD,
      secondary: false,
    },
  },
  {
    path: '/start',
    name: 'Start',
    component: () => import('./components/Pages/Start.vue'),
    meta: {
      icon: PrimeIcons.SERVER,
      secondary: false,
    },
  },
  {
    path: '/',
    name: 'Tracker Features',
    component: () => import('./components/Pages/NoSaves.vue'),
    meta: {
      icon: PrimeIcons.ANDROID,
      secondary: false,
    },
    // beforeCreate(to, from) {
    //   if (localStorage.getItem("saves") != null ) {

    //   }
    // },
    
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

export default router 
// export default routes