import {createRouter, createWebHashHistory} from 'vue-router'
import childTableData from "../view/RightData.vue";
import blank from "../components/AddConnection.vue";

const routes = [
    {
        path: '/',
        name: 'blank',
        component: blank
    },
    {
        path: '/childTable/data',
        name: 'childTableData',
        component: childTableData
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router