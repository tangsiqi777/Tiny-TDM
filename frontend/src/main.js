import {createApp} from 'vue'
import App from './App.vue'
import './assets/css/style.css';
import router from "./util/router.js";
import pinia from './util/myPinia.js'
import "./assets/css/common.css"
import ArcoVueIcon from '@arco-design/web-vue/es/icon';
import ArcoVue, {Message} from '@arco-design/web-vue';
import '@arco-design/web-vue/dist/arco.css';

let app = createApp(App)
Message._context = app._context;
app.use(ArcoVue).use(ArcoVueIcon).use(router).use(pinia).mount('#app')

