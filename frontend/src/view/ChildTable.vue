<script setup>
import {onMounted, ref} from "vue";
import ChildTableList from "../components/ChildTableItem.vue";
import {ListChildTable} from "../../wailsjs/go/main/App.js";
import {Store} from "../store.js";
import {tmitt} from "../mitt.js";
import Back from "../components/Back.vue";
import Search from "../components/Search.vue";
import SqlQuery from "../components/SqlQuery.vue";
import TablePage from "../components/TablePage.vue";

console.log("ChildTable List\n\n\n\n\n\n")
const store = Store()

let childTableList = ref([])


onMounted(() => {
  displayChildTable()
})

tmitt.on("searchChildTable", (data) => {
  console.log(data); //A组件的数据
  displayChildTable()
});

function displayChildTable() {
  let database = store.database
  let superTable = store.superTable
  let childTableSearch = store.childTableSearch
  let start = new Date().getTime();
  ListChildTable(store.conn.conn, database, superTable, childTableSearch).then((childTables) => {
    childTableList.value = childTables
    let end = new Date().getTime();
    console.log("获取子表耗时:" + (end - start) + JSON.stringify(childTables))
  })
}

function toChildTableData(table) {
  console.log("update table")
  store.setChildTable(table)
}


</script>

<template>
  <div class="table-list">
    <div class="search-input">
      <Search :display-type="3" class="search-input1"></Search>
      <SqlQuery></SqlQuery>
    </div>
    <a-scrollbar outer-style="height:calc(100% - 80px);" style="height:100%;overflow: auto;" class="table-item-list">
      <ChildTableList :child-table="item" v-for="item in childTableList" :key="item"
                      @click="toChildTableData(item)"></ChildTableList>
    </a-scrollbar>
    <div class="bottom">
      <Back></Back>
      <TablePage :total="200" class="table-page"></TablePage>
    </div>

  </div>
</template>

<style scoped>
.table-list {
  width: 100%;
  height: 100%;
}

.search-input {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 45px;
  width: 100%;
}

.search-input1 {
  height: 80%;
  width: 70%;
}

.table-item-list {
  height: calc(100% - 80px);
}

.bottom{
  display: flex;
  height: 35px;
  width: 100%;
  align-items: center;
  justify-content: space-between;
}

.table-page{
  width: 200px

}







</style>