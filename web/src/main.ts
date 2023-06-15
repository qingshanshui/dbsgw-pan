import {createApp} from 'vue' // 引入创建vue
import naive from 'naive-ui' // 引入ui框架
import {router} from "@/router"; // 引入router路由
import pinia from "@/stores";// 引入全局状态管理
import './style.css' // 引入全局长沙市
import App from './App.vue' // 引入app.vue入口vue文件

createApp(App).use(pinia).use(naive).use(router).mount('#app') // 实例化vue，挂载插件，绑定元素
