import { createRouter, createWebHashHistory } from 'vue-router'
import Home from './components/Home.vue'
import { useCounterStore } from './stores/counter'
import { useSavesStore } from './stores/saves'
import { PrimeIcons } from 'primevue/api'

// import TrackerFeatures from './components/TrackerFeatures.vue'

const routes = [
  {
    path: '/home',
    name: 'Home',
    component: Home,
    meta: {
      icon: PrimeIcons.HOME,
      secondary: false
    },
    children: [
      {
        path: 'test1',
        name: 'TFeatures',
        component: () => import('./components/TrackerFeatures.vue'),
        meta: {
          secondary: true,
        }
      },
      {
        path: 'test',
        name: 'test',
        component: () => import('./components/Features.vue'),
        meta: {
          secondary: true
        }
      }
    ]
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
    path: '/',
    name: 'Start',
    component: () => import('./components/Start.vue'),
    meta: {
      icon: PrimeIcons.SERVER,
      secondary: false,
    },
  },
  {
    path: '/tracker-features',
    name: 'Tracker Features',
    component: () => import('./components/TrackerFeatures.vue'),
    meta: {
      icon: PrimeIcons.ANDROID,
      secondary: false,
    }
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})



router.beforeEach(() => {
  
  const saves = useSavesStore()
  // if (saves.saves == 1) return '/tracker-features'
})

export default router 
// export default routes