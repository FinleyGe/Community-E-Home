import { createRouter, createWebHistory } from "vue-router";

export default createRouter({
    history: createWebHistory(),
    routes:[
        { path: '/', redirect: '/index'},
        { path: '/index', component: () => import("../views/IndexPage.vue")},
        // { path: '/test', component: () => import("./views/RegisterPage.vue")},
        { path: '/login', name:'login', component: () => import("../views/LoginPage.vue")},
        { path: '/register', name:'register', component: () => import("../views/RegisterPage.vue")},
        { path: '/userinfo', name:'userinfo', component: () => import("../views/UserInfo.vue")},
        // 404 page should be placed in the end of the list of routes.
        { path:'/:pathMatch(.*)*', name:'404', component: () => import("../views/404.vue")}
    ]
})
