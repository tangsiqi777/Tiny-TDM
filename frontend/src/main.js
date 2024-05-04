import {createApp} from 'vue'
import App from './App.vue'
import './style.css';
import router from "./router.js";
import {createPinia} from 'pinia'
import "./assets/css/common.css"
import ArcoVueIcon from '@arco-design/web-vue/es/icon';
import { Message } from '@arco-design/web-vue';


import ArcoVue from '@arco-design/web-vue';
import '@arco-design/web-vue/dist/arco.css';


const pinia = createPinia()
let app = createApp(App)
Message._context = app._context;
app.use(ArcoVue).use(ArcoVueIcon).use(router).use(pinia).mount('#app')

