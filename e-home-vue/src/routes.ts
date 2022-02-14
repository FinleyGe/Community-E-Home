import { createRouter, createWebHistory } from "vue-router";

export default createRouter({
    history: createWebHistory(),
    routes:[
        { path: '/', redirect: '/index'},
        { path: '/index', component: () => import("./views/IndexPage.vue")},
        { path: '/test', component: () => import("./views/IndexPage.vue")},
        { path: '/login', name:'login', component: () => import("./views/RegisterPage.vue")}
    ]
})
