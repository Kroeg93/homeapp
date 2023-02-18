import Home from '../views/Home.vue'
import About from "../views/About.vue";
import {createRouter, createWebHistory} from "vue-router";
import List from "../views/List.vue";
import Lightning from "../views/Lightning/Index.vue"

const routes = [
    {path: '/', name: 'Home', component: Home},
    {path: '/lightning', name: 'Beleuchtung', component: Lightning},
    {path: '/about', name: 'About', component: About},
    {path: '/list', name: 'List', component: List}
]

const router = createRouter({
    history: createWebHistory(),
    routes: routes
})

export default router