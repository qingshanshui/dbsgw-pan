import {createApp} from 'vue'
import naive from 'naive-ui'
import './style.css'
import App from './App.vue'
import {router} from "@/router";
import pinia from "@/stores";

createApp(App).use(pinia).use(naive).use(router).mount('#app')
