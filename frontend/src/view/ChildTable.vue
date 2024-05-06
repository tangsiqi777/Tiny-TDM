<script setup>
import {onMounted, ref} from "vue";
import ChildTableItem from "../components/ChildTableItem.vue";
import {CountChildTable, ListChildTable} from "../../wailsjs/go/service/RestSqlService.js";
import {Store} from "../util/store.js";
import {SingleMitt} from "../util/mitt.js";
import Back from "../components/Back.vue";
import Search from "../components/Search.vue";
import SqlQuery from "../components/SqlQuery.vue";
import {hasError, restDataToJsonObj} from "../version/restDataHandle.js";
import {Message} from "@arco-design/web-vue";

console.log("ChildTable List\n\n\n\n\n\n")
const store = Store()

let childTableList = ref([])

const PAGE_SIZE = 500;

let current = ref(1)
let size = ref(PAGE_SIZE)
let total = ref(0)


onMounted(() => {
  displayChildTable()
})

SingleMitt.on("searchChildTable", (data) => {
  console.log(data); //A组件的数据
  displayChildTable()
});

function changeCurrent(cur) {
  current.value = parseInt(cur)
  displayChildTable();
}

function displayChildTable() {
  let database = store.database
  let superTable = store.superTable
  let childTableSearch = store.childTableSearch
  let start = new Date().getTime();
  CountChildTable(store.conn.conn, database, superTable, childTableSearch).then((countChildTable) => {
    let errorMsg = hasError(countChildTable);
    if (errorMsg !== "") {
      Message.error({
        id: 'countChildTable',
        content: errorMsg,
        duration: 2000
      });
      return;
    }
    total.value = restDataToJsonObj(countChildTable)[0].num
    console.log("count:" + total.value)


    ListChildTable(store.conn.conn, database, superTable, childTableSearch, PAGE_SIZE, current.value).then((childTables) => {
      let errorMsg = hasError(childTables);
      if (errorMsg !== "") {
        Message.error({
          id: 'displayChildTable',
          content: errorMsg,
          duration: 2000
        });
        return;
      }
      childTableList.value = restDataToJsonObj(childTables)
      let end = new Date().getTime();
      console.log("获取子表耗时:" + (end - start) + JSON.stringify(childTables))
    })
  })

}
</script>

<template>
  <div class="table-list">
    <div class="search">
      <Back class="back"></Back>
      <Search :display-type="3" class="search-input"></Search>
      <SqlQuery class="sql-query"></SqlQuery>
    </div>
    <a-scrollbar outer-style="height:calc(100% - 105px);" style="height:100%;overflow: auto;" class="table-item-list">
      <ChildTableItem :child-table="item.tbname" v-for="item in childTableList"
                      :key="item.tbname"></ChildTableItem>
    </a-scrollbar>
    <div class="bottom">
      <a-pagination :total="total" :page-size="size" :current="current" :show-total="true" @change="changeCurrent"
                    simple/>
    </div>
  </div>
</template>

<style scoped>
.table-list {
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

.table-item-list {
  height: calc(100% - 105px);
}

.bottom {
  display: flex;
  height: 60px;
  width: 100%;
  align-items: center;
  justify-content: center;
}
</style>