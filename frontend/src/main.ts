import './assets/main.css'
import 'primeicons/primeicons.css'

import { createApp } from 'vue'
import App from './App.vue'
import {createWebHistory, createRouter} from 'vue-router'
import VueApexCharts from "vue3-apexcharts";

import HomeView from './views/HomeView.vue'
import Monitoring from './views/Monitoring.vue'
import Auth from './views/Auth.vue'

const routes = [
    { path: '/', component: HomeView },
    { path: '/auth', component: Auth },
    { path: '/monitoring', component: Monitoring },
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

const app = createApp(App)
app.use(router)
app.use(VueApexCharts)
app.mount('#app')