import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from './router'
import {library} from '@fortawesome/fontawesome-svg-core'
import {faLightbulb} from '@fortawesome/free-solid-svg-icons'
import '@fortawesome/free-brands-svg-icons'
import '@fortawesome/free-regular-svg-icons'

import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

const vuetify = createVuetify({
    components,
    directives,
})


library.add(faLightbulb)

createApp(App)
    .use(router)
    .use(vuetify)
    .mount('#app')
