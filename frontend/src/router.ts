import { createRouter, createWebHashHistory } from 'vue-router'
import Home from './components/Home.vue'
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
      icon: PrimeIcons.HOME,
      secondary: false
    },
    
  },
  {
    path: '/features',
    name: 'Features',
    component: () => import('./components/Features.vue'),
    meta: {
      icon: PrimeIcons.DISCORD,
      secondary: false,
    },
  },
  {
    path: '/start',
    name: 'Start',
    component: () => import('./components/Start.vue'),
    meta: {
      icon: PrimeIcons.SERVER,
      secondary: false,
    },
  },
  {
    path: '/',
    name: 'Tracker Features',
    component: () => import('./components/TrackerFeatures.vue'),
    meta: {
      icon: PrimeIcons.ANDROID,
      secondary: false,
    },
    
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})



// router.beforeEach((to, from, next) => {
//   // const saves = localStorage.getItem('saveCount')
//   if (localStorage.getItem('saves') == "0") next({ name: 'Tracker Features'})
//   else next()
  
  
//   // if (saves.saves == 1) return '/tracker-features'
// })


export default router 
// export default routes