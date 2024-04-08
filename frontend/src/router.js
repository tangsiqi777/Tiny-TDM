import {createRouter, createWebHashHistory, createWebHistory} from 'vue-router'
import database from './view/Database.vue'
import superTable from './view/SuperTable.vue'
import childTable from './view/ChildTable.vue'
import childTableData from "./view/RightData.vue";
import blank from "./view/Blank.vue";

const routes = [
    // {
    //     path: '/',
    //     name: 'server',
    //     component: server
    // },
    {
        path: '/',
        name: 'database',
        components: {
            default: database,
            rightData: blank
        }
    },
    {
        path: '/superTable',
        name: 'superTable',
        components: {
            default: superTable,
            rightData: blank
        }
    },
    {
        path: '/childTable',
        name: 'childTable',
        components: {
            default: childTable,
            rightData: blank
        }
    },
    {
        path: '/childTable/data',
        name: 'childTableData',
        components: {
            default: childTable,
            rightData: childTableData
        }
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router