import { RouteLocationNormalized, createRouter, createWebHashHistory } from 'vue-router'
import Home from './components/Pages/DefaultHome.vue'
import { useCounterStore } from './stores/counter'
import { useSavesStore } from './stores/saves'
import { PrimeIcons } from 'primevue/api'
import { GetNumSaves } from '../wailsjs/go/main/App'

const routes = [
  {
    path: '/',
    name: 'App Home',
    redirect: {path: '/home'},
    component: () => import('./components/Pages/Home.vue'),
    meta: {
      hidden: true,
      totals: false,
    },
    beforeEnter: async (to: RouteLocationNormalized) => {
      const numSaves: number = await GetNumSaves();
      if (numSaves == 0) {
        localStorage.setItem("saves", "0");
        localStorage.setItem("defaultSave", "");
        return;
      }
      const dSave: string | null = localStorage.getItem("defaultSave")
      if (dSave != null) {
        if (dSave != "0" && dSave != "") {
          console.log('here')
          return {path: "/save/" + dSave + "/home", replace: true}
        }
      }
    },
    children: [
      {
        path: 'home',
        name: 'All Saves Home',
        component: () => import('./components/Pages/DefaultHome.vue'),
        meta: {
          hidden: true,
          totals: true
        }
      },
      {
        path: 'results',
        name: 'All Results',
        component: () => import('./components/Pages/AllResults.vue'),
        meta: {
          hidden: true,
          totals: true,
        }
      },
      {
        path: 'transfers',
        name: 'All Transfers',
        component: () => import('./components/Pages/AllTransfers.vue'),
        meta: {
          hidden: true,
          totals: true,
        }
      },
      {
        path: 'player-seasons',
        name: 'All Saves Seasons',
        component: () => import('./components/Pages/AllIndivSquadSeasons.vue'),
        meta: {
          hidden: true,
          totals: true
        }
      },
      {
        path: 'player-totals',
        name: 'All Saves Totals',
        component: () => import('./components/Pages/AllSquadTotals.vue'),
        meta: {
          hidden: true,
          totals: true,
        }
      }
    ]
  },
  {
    path: '/no-saves',
    name: 'No Saves',
    component: () => import('./components/Pages/NoSaves.vue'),
    meta: {
      icon: PrimeIcons.ANDROID,
      hidden: true,
      totals: false,
    },
  },
  {
    path: '/save/:id',
    name: 'Save',
    component: () => import('./components/Pages/Save.vue'),
    meta: {
      hidden: true,
      totals: false,
    },

    children: [
      {
        path: 'home',
        name: 'Home',
        component: () => import('./components/Pages/SaveHome.vue'),
        meta: {
          hidden: false,
          totals: false
        },
        beforeEnter: (to: RouteLocationNormalized, from: RouteLocationNormalized) => {
          if ("id" in to.params) {
            if (to.params.id == "0") {
              return { name: "App Home", replace: true}
            }
          }
        }
      },
      {
        path: 'results',
        name: 'Results',
        component: () => import('./components/Pages/Results.vue'),
        meta: {
          hidden: false,
          totals: false
        }
      },
      {
        path: 'transfers',
        name: 'Transfers',
        component: () => import('./components/Pages/Transfers.vue'),
        meta: {
          hidden: false,
          totals: false
        }
      },
      {
        path: 'player-seasons',
        name: 'Player Seasons',
        component: () => import('./components/Pages/IndivSquadSeasons.vue'),
        meta: {
          hidden: false,
          totals: false
        }
      },
      {
        path: 'player-totals',
        name: 'Player Totals',
        component: () => import('./components/Pages/SquadTotals.vue'),
        meta: {
          hidden: false,
          totals: false
        }
      }
    ]
  },
  {
    path: '/:catchAll(.*)',
    redirect: '/'
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes,
})


export default router 