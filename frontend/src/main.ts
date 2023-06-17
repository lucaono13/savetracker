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
import InputMask from 'primevue/inputmask'
import Divider from 'primevue/divider'
import Textarea from 'primevue/textarea'
import Toast from 'primevue/toast'
import ToastService from 'primevue/toastservice'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import Badge from 'primevue/badge'
// import ColumnGroup from 'primevue/columngroup'
// import Row from 'primevue/row'
import MultiSelect from 'primevue/multiselect'
import TabView from 'primevue/tabview'
import TabPanel from 'primevue/tabpanel'
import AutoComplete from 'primevue/autocomplete'
import Carousel from 'primevue/carousel'
import Fieldset from 'primevue/fieldset'
import DataView from 'primevue/dataview';
import DataViewLayoutOptions from 'primevue/dataviewlayoutoptions'   // optional
import InlineMessage from 'primevue/inlinemessage'
import Tooltip from 'primevue/tooltip'
import Editor from 'primevue/editor'
import Chart from 'primevue/chart'
import InputSwitch from 'primevue/inputswitch'

// Flags
// @ts-ignore
// import Flag from 'vue-flagpack'

// Font Awesome Icons
//import the fontawesome core
import { library } from '@fortawesome/fontawesome-svg-core'

// Import Font Awesome icon component
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

import { config } from "@fortawesome/fontawesome-svg-core"

import { fas } from "@fortawesome/free-solid-svg-icons"

// Import specific icons
import { faMugHot, faList, faQuestion, faHouseChimney } from '@fortawesome/free-solid-svg-icons'

// Flags

// import Flag from 'vue-flagpack'

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
app.use(ToastService)
// app.use(Flag, { name: 'Flag'})

// Flags
import CountryFlag from 'vue-country-flag-next'
app.component('CountryFlag', CountryFlag)

// const app = createApp(App).use(PrimeVue).use(store, key).use(router)//.mount('#app')

// Add Icons to Library
// library.add(faMugHot)
// library.add(faList)
// library.add(faQuestion)
library.add(fas)

// Add FontAwesome to app
app.component('font-awesome-icon', FontAwesomeIcon)

config.styleDefault = "solid"


// Adding PrimeVue
app.use(PrimeVue)

// PrimeVue components
app.component('ProgressSpinner',ProgressSpinner)
    .component('ProgressBar', ProgressBar)
    .component('Sidebar',Sidebar)
    .component('Menubar', Menubar)
    .component('InputText', InputText)
    .component('MegaMenu', MegaMenu)
    .component('Menu', Menu)
    .component('Dock', Dock)
    .component('Button', Button)
    .component('Dialog', Dialog)
    .component('Dropdown', Dropdown)
    .component('Card', Card)
    .component('BlockUI', BlockUI)
    .component('OverlayPanel', OverlayPanel)
    .component('CascadeSelect', CascadeSelect)
    .component('VirtualScroller', VirtualScroller)
    .component('Checkbox', Checkbox)
    .component('Image', Image)
    .component('InputMask', InputMask)
    .component('Divider', Divider)
    .component('Textarea', Textarea)
    .component('Toast', Toast)
    .component('DataTable', DataTable)
    .component('Column', Column)
    .component('Badge', Badge)
    .component('MultiSelect', MultiSelect)
    .component('TabView', TabView)
    .component('TabPanel', TabPanel)
    .component('AutoComplete', AutoComplete)
    .component('Carousel', Carousel)
    .component('Fieldset', Fieldset)
    .component('DataView', DataView)
    .component('DataViewLayoutOptions', DataViewLayoutOptions)
    .component('InlineMessage', InlineMessage)
    .component('Editor', Editor)
    .component('Chart',Chart)
    .component('InputSwitch', InputSwitch)
    .directive('tooltip', Tooltip)

// app.use(Flag)
// Use Vue Router
app.use(router)

// Mounting app
app.mount('#app')