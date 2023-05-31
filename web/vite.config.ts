import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import {resolve} from 'path';

const pathResolve = (dir: string): any => {
    return resolve(__dirname, '.', dir);
};
const alias: Record<string, string> = {
    '/@': pathResolve('./src/'),
    'vue-i18n': 'vue-i18n/dist/vue-i18n.cjs.js',
};

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue()],
    resolve: {alias},
    server: {
        proxy: {
            '/v1': 'http://127.0.0.1:8005', // 本地接口
        },
    }

})
