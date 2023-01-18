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
    beforeEnter() {
      const saves = localStorage.getItem("saves")
      const defaultSave = localStorage.getItem("defaultSave")
      if (saves !== null && defaultSave != null) {
        // let defaultSave: string | null = localStorage.getItem("defaultSave")
        // console.log("ok", defaultSave?.length)
        // if (defaultSave?.length !== 0) {
        //   console.log('what')
        // }
        console.log(defaultSave.length !== 0)
        if (defaultSave.length !== 0) {
          console.log('default')
          return { path: '/save/' + defaultSave, replace: true}
        } else {
          console.log('no default')
        }
      }
    },
  },
  {
    path: '/save/:id',
    name: 'save',
    component: () => import('./components/Pages/SaveHome.vue'),
    meta: {
      secondary: false,
    },
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})

export default router 
// export default routes