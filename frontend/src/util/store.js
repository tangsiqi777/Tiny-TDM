import {defineStore} from 'pinia'
import {reactive, ref} from 'vue'

// 用户模块 token setToken removeToken
export const Store = defineStore('store', () => {

    const displayType = ref(0)
    const setDisplayType = (newDisplayType) => {
        displayType.value = newDisplayType
        removeSuperTableSearch()
        removeChildTableSearch()
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

    // 从哪个超级表点击进去的
    const superTable = ref('')
    const setSuperTable = (newSuperTable) => {
        superTable.value = newSuperTable
    }
    const removeSuperTable = () => {
        superTable.value = ''
    }

    //选中查询超级表
    const selectedSuperTable = ref('')
    const setSelectedSuperTable = (superTable) => {
        selectedSuperTable.value = superTable
    }
    const removeSelectedSuperTable = () => {
        selectedSuperTable.value = ''
    }

    const childTable = ref('')
    const setChildTable = (newChildTable) => {
        childTable.value = newChildTable
    }
    const removeChildTable = () => {
        childTable.value = ''
    }

    const superTableSearch = ref('')
    const setSuperTableSearch = (searchKey) => {

        superTableSearch.value = searchKey
    }
    const removeSuperTableSearch = () => {
        superTableSearch.value = ''
    }


    const childTableSearch = ref('')
    const setChildTableSearch = (searchKey) => {
        childTableSearch.value = searchKey
    }
    const removeChildTableSearch = () => {
        childTableSearch.value = ''
    }

    const primaryId = ref('')
    const setPrimaryId = (primary) => {
        primaryId.value = primary
    }
    const removePrimaryId = () => {
        primaryId.value = ''
    }

    const sqlQuerySelected = ref(false)
    const setSqlQuerySelected = () => {
        sqlQuerySelected.value = true
    }
    const removeSqlQuerySelected = () => {
        sqlQuerySelected.value = false
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
        removeDisplayType,
        childTableSearch,
        setChildTableSearch,
        removeChildTableSearch,
        superTableSearch,
        setSuperTableSearch,
        removeSuperTableSearch,
        primaryId,
        setPrimaryId,
        removePrimaryId,
        sqlQuerySelected,
        setSqlQuerySelected,
        removeSqlQuerySelected,
        selectedSuperTable,
        setSelectedSuperTable,
        removeSelectedSuperTable
    }
})