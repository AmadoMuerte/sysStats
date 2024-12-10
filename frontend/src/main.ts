import './assets/main.css'
import 'primeicons/primeicons.css'

import { createApp } from 'vue'
import App from './App.vue'
import {createWebHistory, createRouter} from 'vue-router'
import VueApexCharts from "vue3-apexcharts";

import HomeView from './views/HomeView.vue'
import Monitoring from './views/Monitoring.vue'
import Login from './views/Login.vue'

const routes = [
    { path: '/', component: HomeView },
    { path: '/login', component: Login },
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