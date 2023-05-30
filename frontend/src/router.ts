import { createRouter, createWebHashHistory } from 'vue-router'
import Home from './components/Pages/Home.vue'
import { useCounterStore } from './stores/counter'
import { useSavesStore } from './stores/saves'
import { PrimeIcons } from 'primevue/api'

// import TrackerFeatures from './components/TrackerFeatures.vue'

const routes = [
  {
    path: '/noDefault',
    name: 'No Default',
    component: () => import('./components/Pages/NoDefaultSave.vue'),
    meta: {
      secondary: true,
    }
  },
  {
    path: '/',
    name: 'No Saves',
    component: () => import('./components/Pages/NoSaves.vue'),
    meta: {
      icon: PrimeIcons.ANDROID,
      secondary: true,
    },
    beforeEnter() {
      const saves: string | null = localStorage.getItem("saves")
      const defaultSave: string | null = localStorage.getItem("defaultSave")
      if (saves === null) {
        localStorage.setItem("saves", "0")
      }
      if (saves !== null && defaultSave == null) {
        return { path: 'No Default', replace: true}
      }
      if (saves !== null && defaultSave != null) {
        if (defaultSave.length !== 0) {
          if (defaultSave == "0") {
            return
          }
          return { path: '/save/' + defaultSave + '/home', replace: true}
        }
      }
    },
  },
  {
    path: '/save/:id',
    name: 'Save',
    component: () => import('./components/Pages/Save.vue'),
    meta: {
      secondary: true,

    },
    children: [
      {
        path: 'home',
        name: 'Home',
        component: () => import('./components/Pages/SaveHome.vue'),
        meta: {
          secondary: false,
        }
      },
      {
        path: 'results',
        name: 'Results',
        component: () => import('./components/Pages/Results.vue'),
        meta: {
          secondary: false,
        }
      },
      {
        path: 'transfers',
        name: 'Transfers',
        component: () => import('./components/Pages/Transfers.vue'),
        meta: {
          secondary: false,
        }
      }
    ]
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

export default router 
// export default routes