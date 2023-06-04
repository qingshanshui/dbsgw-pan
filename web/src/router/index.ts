import { createRouter, createWebHashHistory, createWebHistory } from 'vue-router';

export const notFoundAndNoPower = [
    {
        path: '/login',
        name: "登录",
        component: () => import('@/views/login/index.vue'),
    },
    {
        path: '/:path(.*)*',
        name: "网盘",
        component: () => import('@/views/home/index.vue'),
    },
];

// 创建 路由实例
export const router = createRouter({
    // history: createWebHashHistory(), // 本地调试
    history: createWebHistory(), // 线上  加nginx代理
    routes: [...notFoundAndNoPower],
});