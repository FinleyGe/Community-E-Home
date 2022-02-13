import { createApp } from 'vue'
import App from './App.vue'
import router from './routes'
import 'animate.css/animate.min.css'
createApp(App).
use(router).
mount('#app')
