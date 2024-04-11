import {defineStore} from 'pinia'
import {reactive, ref} from 'vue'

// 用户模块 token setToken removeToken
export const Store = defineStore('big-user', () => {

    const displayType = ref(0)
    const setDisplayType = (newDisplayType) => {
        displayType.value = newDisplayType
    }
    const removeDisplayType = () => {
        displayType.value = 0
    }

    const conn = reactive({
        conn: {}
    })
    const setConn = (newConn) => {
        conn.conn = newConn
    }
    const removeConn = () => {
        conn.conn = {}
    }

    const database = ref('')
    const setDatabase = (newDatabase) => {
        database.value = newDatabase
    }
    const removeDatabase = () => {
        database.value = ''
    }

    const superTable = ref('')
    const setSuperTable = (newSuperTable) => {
        superTable.value = newSuperTable
    }
    const removeSuperTable = () => {
        superTable.value = ''
    }

    const childTable = ref('')
    const setChildTable = (newChildTable) => {
        childTable.value = newChildTable
    }
    const removeChildTable = () => {
        childTable.value = ''
    }

    return {
        conn,
        setConn,
        removeConn,
        database,
        setDatabase,
        removeDatabase,
        superTable,
        setSuperTable,
        removeSuperTable,
        childTable,
        setChildTable,
        removeChildTable,
        displayType,
        setDisplayType,
        removeDisplayType
    }
})