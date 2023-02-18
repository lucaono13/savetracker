import { createApp, defineComponent } from 'vue'
import App from './App.vue'
import { useCounterStore, store } from './stores/counter'
import router from './router'
import { createPinia, storeToRefs } from 'pinia'

// PrimeVue Components 
import PrimeVue from 'primevue/config'
import ProgressSpinner from 'primevue/progressspinner'
import ProgressBar from 'primevue/progressbar'
import Sidebar from 'primevue/sidebar'
import Menubar from 'primevue/menubar'
import InputText from 'primevue/inputtext'
import MegaMenu from 'primevue/megamenu'
import Menu from 'primevue/menu'
import Dock from 'primevue/dock'
import Button from 'primevue/button'
import Dialog from 'primevue/dialog'
import Checkbox from 'primevue/checkbox'
import Dropdown from 'primevue/dropdown'
import Card from 'primevue/card'
import BlockUI from 'primevue/blockui'
import OverlayPanel from 'primevue/overlaypanel'
import CascadeSelect from 'primevue/cascadeselect'
import VirtualScroller from 'primevue/virtualscroller'
import Image from 'primevue/image'

// Font Awesome Icons
//import the fontawesome core
import { library } from '@fortawesome/fontawesome-svg-core'

// Import Font Awesome icon component
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

// Import specific icons
import { faMugHot, faList } from '@fortawesome/free-solid-svg-icons'



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

// Add Icons to Library
library.add(faMugHot)
library.add(faList)

// Add FontAwesome to app
app.component('font-awesome-icon', FontAwesomeIcon)

// Adding PrimeVue
app.use(PrimeVue)

// PrimeVue components
app.component('ProgressSpinner',ProgressSpinner)
app.component('ProgressBar', ProgressBar)
app.component('Sidebar',Sidebar)
app.component('Menubar', Menubar)
app.component('InputText', InputText)
app.component('MegaMenu', MegaMenu)
app.component('Menu', Menu)
app.component('Dock', Dock)
app.component('Button', Button)
app.component('Dialog', Dialog)
app.component('Dropdown', Dropdown)
app.component('Card', Card)
app.component('BlockUI', BlockUI)
app.component('OverlayPanel', OverlayPanel)
app.component('CascadeSelect', CascadeSelect)
app.component('VirtualScroller', VirtualScroller)
app.component('Checkbox', Checkbox)
    .component('Image', Image)

// Use Vue Router
app.use(router)

// Mounting app
app.mount('#app')