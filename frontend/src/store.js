import {defineStore} from 'pinia'
import {ref} from 'vue'

// 用户模块 token setToken removeToken
export const Store = defineStore('big-user', () => {

    const displayType = ref(1)
    const setDisplayType = (newDisplayType) => {
        displayType.value = newDisplayType
    }
    const removeDisplayType = () => {
        displayType.value = 1
    }

    const conn = ref('')
    const setConn = (newConn) => {
        conn.value = newConn
    }
    const removeConn = () => {
        conn.value = ''
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