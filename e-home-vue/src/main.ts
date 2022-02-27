import { createApp } from 'vue'
import App from './App.vue'
import router from './routes'
import {store, key } from './store'
import 'animate.css/animate.min.css'
createApp(App).
use(router).
use(store, key).
mount('#app')
