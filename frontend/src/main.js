import {createApp} from 'vue'
import App from './App.vue'
import './style.css';
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import router from "./router.js";
import {createPinia} from 'pinia'


const pinia = createPinia()
createApp(App).use(ElementPlus).use(router).use(pinia).mount('#app')
