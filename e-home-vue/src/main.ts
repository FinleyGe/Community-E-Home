import { createApp } from 'vue'
import App from './App.vue'
import router from './routes/routes'
import 'animate.css/animate.min.css'
import { createPinia } from 'pinia'

createApp(App).
use(router).
// use(store, key).
use(createPinia()).
mount('#app')