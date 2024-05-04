<script setup>

import SuperTableItem from "../components/SuperTableItem.vue";
import {ref} from "vue";
import {DescTable, ListSuperTable} from "../../wailsjs/go/service/RestSqlService.js";
import {Store} from "../store.js";
import {SingleMitt} from "../mitt.js";
import Search from "../components/Search.vue";
import Back from "../components/Back.vue";
import SqlQuery from "../components/SqlQuery.vue";
import {parseNestedJsonAndGetData} from "../TdengineRestData.js";


console.log("SuperTable List\n\n\n\n\n\n")

let superTableList = ref([])

const store = Store()


SingleMitt.on("searchSuperTable", (data) => {
  displaySuperTable()
});

displaySuperTable()

function displaySuperTable() {
  let database = store.database
  let superTableSearch = store.superTableSearch
  ListSuperTable(store.conn.conn, database, superTableSearch).then((superTables) => {
    superTableList.value = parseNestedJsonAndGetData(superTables)
  })
}

function toChildTable(superTable) {
  let database = store.database
  DescTable(store.conn.conn, database, superTable).then((tableInfoRet) => {
    let tableInfo = parseNestedJsonAndGetData(tableInfoRet)
    console.log("tableInfo:" + JSON.stringify(tableInfo) )
    if (tableInfo !== undefined && tableInfo !== null && tableInfo.length > 0) {
      console.log("tableInfo:" + tableInfo[0].field)
      store.setPrimaryId(tableInfo[0].field)
    }
    store.setSuperTable(superTable)
    store.setDisplayType(3)
  })

}

</script>


<template>
  <div class="super-table">
    <div class="search">
      <Back class="back"></Back>
      <Search :display-type="2" class="search-input"></Search>
      <SqlQuery class="sql-query"></SqlQuery>
    </div>

    <a-scrollbar outer-style="height:calc(100% - 45px);" style="height:100%;overflow: auto;" class="table-list">
      <SuperTableItem :superTable="item.stableName" v-for="item in superTableList" :key="item"
                      @click="toChildTable(item.stableName)"></SuperTableItem>
    </a-scrollbar>
  </div>
</template>

<style scoped>
.super-table {
  width: 100%;
  height: 100%;
}

.search {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 45px;
  width: 100%;
}

.search-input {
  height: 70%;
  width: 70%;
}

.back {
  width: 15%;
  height: 45px;
}

.sql-query {
  height: 45px;
  width: 15%;
}

.table-list {
  height: calc(100% - 45px);
}


</style>