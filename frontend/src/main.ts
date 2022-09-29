import { createApp, defineComponent } from 'vue'
import App from './App.vue'
import { useCounterStore, store } from './stores/counter'
import router from './router'
import { createPinia, storeToRefs } from 'pinia'

// PrimeVue Components 
import PrimeVue from 'primevue/config'
import ProgressSpinner from 'primevue/progressspinner'
import Sidebar from 'primevue/sidebar'
import Menubar from 'primevue/menubar'
import InputText from 'primevue/inputtext'
import MegaMenu from 'primevue/megamenu'
import Menu from 'primevue/menu'
import Dock from 'primevue/dock'
import Button from 'primevue/button'

// PrimeIcons
import 'primeicons/primeicons.css'

// PrimeFlex
import 'primeflex/primeflex.css'

// Themeing
import 'primevue/resources/primevue.min.css'
import 'primevue/resources/themes/lara-dark-teal/theme.css'

// const pinia = createPinia()

const app = createApp(App)
app.use(createPinia())

// const app = createApp(App).use(PrimeVue).use(store, key).use(router)//.mount('#app')

// Adding PrimeVue
app.use(PrimeVue)

// PrimeVue components
app.component('ProgressSpinner',ProgressSpinner)
app.component('Sidebar',Sidebar)
app.component('Menubar', Menubar)
app.component('InputText', InputText)
app.component('MegaMenu', MegaMenu)
app.component('Menu', Menu)
app.component('Dock', Dock)
app.component('Button', Button)

// Use Vue Router
app.use(router)

// Mounting app
app.mount('#app')