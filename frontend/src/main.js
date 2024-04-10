import {createApp} from 'vue'
import App from './App.vue'
import './style.css';
import 'element-plus/dist/index.css'
import router from "./router.js";
import {createPinia} from 'pinia'


import ArcoVue from '@arco-design/web-vue';
import '@arco-design/web-vue/dist/arco.css';


const pinia = createPinia()
createApp(App).use(ArcoVue).use(router).use(pinia).mount('#app')
