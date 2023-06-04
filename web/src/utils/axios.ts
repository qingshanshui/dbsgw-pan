import axios from 'axios';
// 配置新建一个 axios 实例
const service = axios.create({
    baseURL: '/',
    timeout: 50000,
    headers: {'Content-Type': 'application/json'},
});

// 添加请求拦截器
service.interceptors.request.use(
    (config) => {
        // 在发送请求之前做些什么 token
        if (localStorage.getItem('token')) {
            config.headers['Authorization'] = `${localStorage.getItem('token')}`;
        }
        console.log("请求拦截：", config)
        return config;
    },
    (error) => {
        // 对请求错误做些什么
        return Promise.reject(error);
    }
);

// 添加响应拦截器
service.interceptors.response.use(
    (response) => {
        console.log("响应拦截：", response)
        return response
    },
    (error) => {
        // 对响应错误做些什么
        return Promise.reject(error);
    }
);

// 导出 axios 实例
export default service;
