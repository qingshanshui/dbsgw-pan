import {createApp} from 'vue'
import naive from 'naive-ui'
import {router} from "@/router";
import pinia from "@/stores";
import './style.css'
import App from './App.vue'

createApp(App).use(pinia).use(naive).use(router).mount('#app')
