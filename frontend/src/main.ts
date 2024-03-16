import { createApp } from 'vue'
import { createRouter, createWebHistory } from 'vue-router'
import '@/assets/style.css'
import '@/assets/index.css'
import App from '@/App.vue'

import HomeView from '@/views/HomeView.vue'
import AboutView from '@/views/AboutView.vue'
import ContactView from '@/views/ContactView.vue'
import DashboardView from '@/views/DashboardView.vue'
import StatisticsView from '@/views/StatisticsView.vue'
import GoogleMapLoader from '@/components/GoogleMapLoader.vue'
import AuthView from '@/views/AuthView.vue'

import { isUserLoggedIn } from '@/lib/utils'

const app = createApp(App);

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: '/', component: HomeView },
        { path: '/about', component: AboutView },
        { path: '/contact', component: ContactView },
        {
            path: '/login',
            component: AuthView,
            beforeEnter: (_, __, next) => {
                if (isUserLoggedIn()) {
                    next('/');
                } else {
                    next();
                }
            },
        },
        {
            path: '/signup',
            component: AuthView,
            beforeEnter: (_, __, next) => {
                if (isUserLoggedIn()) {
                    next('/');
                } else {
                    next();
                }
            },
        },
        { 
            path: '/dashboard',
            component: DashboardView,
            children: [
                { path: '', component: GoogleMapLoader, name: 'Map' },
                { path: 'statistics', component: StatisticsView, name: 'Statistics' },
            ]
        },
        // default redirect to home page
        { path: '/:pathMatch(.*)*', redirect: '/' }
    ],
});

app.use(router);
app.mount('#app')
