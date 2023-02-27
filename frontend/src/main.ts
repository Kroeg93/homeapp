import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import './index.css'
import router from './router'
import {library} from '@fortawesome/fontawesome-svg-core'
import {faLightbulb} from '@fortawesome/free-solid-svg-icons'
import {FontAwesomeIcon} from "@fortawesome/vue-fontawesome";
import '@fortawesome/free-brands-svg-icons'
import '@fortawesome/free-regular-svg-icons'

library.add(faLightbulb)

createApp(App)
    .use(router)
    .mount('#app')
