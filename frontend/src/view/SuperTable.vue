<script setup>

import SuperTableItem from "../components/SuperTableItem.vue";
import {ref} from "vue";
import {ListSuperTable} from "../../wailsjs/go/service/RestSqlService.js";
import {Store} from "../util/store.js";
import {SingleMitt} from "../util/mitt.js";
import Search from "../components/Search.vue";
import Back from "../components/Back.vue";
import SqlQuery from "../components/SqlQuery.vue";
import {hasError, restDataToJsonObj} from "../version/restDataHandle.js";
import {Message} from "@arco-design/web-vue";


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
    let errorMsg = hasError(superTables);
    if (errorMsg !== "") {
      Message.error({
        id: 'displaySuperTable',
        content: errorMsg,
        duration: 2000
      });
      return;
    }
    superTableList.value = restDataToJsonObj(superTables)
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
      <SuperTableItem :superTable="item.stable_name" v-for="item in superTableList" :key="item"></SuperTableItem>
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
  height: 72%;
  width: calc(100% - 90px);
}

.back {
  width: 45px;
  height: 45px;
}

.sql-query {
  height: 45px;
  width: 45px;
}

.table-list {
  height: calc(100% - 45px);
}


</style>