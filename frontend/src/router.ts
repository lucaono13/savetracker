import { RouteLocationNormalized, createRouter, createWebHashHistory } from 'vue-router'
import Home from './components/Pages/Home.vue'
import { useCounterStore } from './stores/counter'
import { useSavesStore } from './stores/saves'
import { PrimeIcons } from 'primevue/api'
import { GetNumSaves, LogPls } from '../wailsjs/go/main/App'

// import TrackerFeatures from './components/TrackerFeatures.vue'

const routes = [
  {
    path: '/noDefault',
    name: 'No Default',
    component: () => import('./components/Pages/NoDefaultSave.vue'),
    meta: {
      secondary: true,
    },
    // beforeEnter: (to: RouteLocationNormalized) => {
    //   console.log(to)
    // }
  },
  {
    path: '/',
    name: 'No Saves',
    component: () => import('./components/Pages/NoSaves.vue'),
    meta: {
      icon: PrimeIcons.ANDROID,
      secondary: true,
    },
    beforeEnter: async (to: RouteLocationNormalized) => {
      // ...
      // if (to.name == "No Default" || to.name == "No Saves") {
      //   return;
      // }
      const numSaves: number = await GetNumSaves();

      if (numSaves == 0) {
        localStorage.setItem("saves", "0");
        localStorage.setItem("defaultSave", "");
        return { name: "No Saves", replace: true };
      }
      // LogPls(("params" in to).toString())
      // console.log(to)
      const dSave: string | null = localStorage.getItem("defaultSave")
      console.log(dSave)
      if (dSave != null) {
        if (dSave != "0") {
          return {path: "/save/" + dSave + "/home", replace: true}
        }
      } else {
        return {name: "No Default"}
      }
    }
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
        },
        beforeEnter: (to: RouteLocationNormalized) => {
          if ("id" in to.params) {
            if (to.params.id == "0") {
              return { name: "No Default", replace: true}
            }
          }
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
      },
      {
        path: 'playerSeasons',
        name: 'Player Seasons',
        component: () => import('./components/Pages/IndivSquadSeasons.vue'),
        meta: {
          secondary: false,
        }
      },
      {
        path: 'playerTotals',
        name: 'Player Totals',
        component: () => import('./components/Pages/SquadTotals.vue'),
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