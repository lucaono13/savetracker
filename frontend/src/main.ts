import { createApp } from 'vue'
import App from './App.vue'
import { store, key } from './store'
import router from './router'

// PrimeVue
import PrimeVue from 'primevue/config'
import ProgressSpinner from 'primevue/progressspinner'

// Quasar
import { Quasar } from 'quasar'

// Import icon libraries
import '@quasar/extras/material-icons/material-icons.css'

// Import Quasar css
import 'quasar/src/css/index.sass'

const app = createApp(App)

// const app = createApp(App).use(PrimeVue).use(store, key).use(router)//.mount('#app')

// // Adding PrimeVue
// app.use(PrimeVue)
// // PrimeVue components
// app.component('ProgressSpinner',ProgressSpinner)

app.use(Quasar, {
    plugins: {},
})

// For Wails
app.use(store, key).use(router)

// Mounting app
app.mount('#app')

