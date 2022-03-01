import { createRouter, createWebHistory } from "vue-router";

export default createRouter({
    history: createWebHistory(),
    routes:[
        { path: '/', redirect: '/index'},
        { path: '/index', component: () => import("../views/IndexPage.vue")},
        // { path: '/test', component: () => import("./views/RegisterPage.vue")},
        { path: '/login', name:'login', component: () => import("../views/LoginPage.vue")},
        { path: '/register', name:'register', component: () => import("../views/RegisterPage.vue")},
        { path:'/:pathMatch(.*)*', name:'404', component: () => import("../views/404.vue")}
    ]
})
