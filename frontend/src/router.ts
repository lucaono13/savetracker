import { createRouter, createWebHashHistory } from 'vue-router'
import Home from './components/Pages/Home.vue'
import { useCounterStore } from './stores/counter'
import { useSavesStore } from './stores/saves'
import { PrimeIcons } from 'primevue/api'

// import TrackerFeatures from './components/TrackerFeatures.vue'

const routes = [
  // {
  //   path: '/home',
  //   name: 'Home 2',
  //   component: Home,
  //   meta: {
  //     icon: "fa-solid fa-mug-hot",
  //     secondary: false
  //   },
    
  // },
  // {
  //   path: '/features',
  //   name: 'Features',
  //   component: () => import('./components/Pages/Features.vue'),
  //   meta: {
  //     icon: PrimeIcons.DISCORD,
  //     secondary: false,
  //   },
  // },
  // {
  //   path: '/start',
  //   name: 'Start',
  //   component: () => import('./components/Pages/Start.vue'),
  //   meta: {
  //     icon: PrimeIcons.SERVER,
  //     secondary: false,
  //   },
  // },
  {
    path: '/noDefault',
    name: 'No Default',
    component: () => import('./components/Pages/NoDefaultSave.vue'),
    meta: {
      secondary: false,
    }
  },
  {
    path: '/',
    name: 'No Saves',
    component: () => import('./components/Pages/NoSaves.vue'),
    meta: {
      icon: PrimeIcons.ANDROID,
      secondary: false,
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
          return { path: '/save/' + defaultSave, replace: true}
        }
      }
    },
  },
  {
    path: '/save/:id',
    name: 'Save Home',
    component: () => import('./components/Pages/SaveHome.vue'),
    meta: {
      secondary: false,

    },
    // children: [
    //   {
    //     path: 'home',
    //     name: 'Save Home',
    //     component: () => import('./components/Pages/SaveHome.vue'),
    //     meta: {
    //       secondary: false,
    //     }
    //   }
    // ]
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

export default router 
// export default routes