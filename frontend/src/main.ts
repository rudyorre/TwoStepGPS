import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import '@/assets/style.css'
import '@/assets/index.css'
import App from '@/App.vue'

import HomeView from '@/views/HomeView.vue'
import AboutView from '@/views/AboutView.vue'
import ContactView from '@/views/ContactView.vue'
import DashboardView from '@/views/DashboardView.vue'

const app = createApp(App);

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: '/', component: HomeView },
        { path: '/about', component: AboutView },
        { path: '/contact', component: ContactView },
        { path: '/dashboard', component: DashboardView },
    ],
});

app.use(router);
app.mount('#app')
