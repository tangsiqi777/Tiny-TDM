<script setup>
import {onMounted, ref} from "vue";
import ChildTableList from "../components/ChildTableItem.vue";
import {ListChildTable} from "../../wailsjs/go/main/App.js";
import {Store} from "../store.js";
import {tmitt} from "../mitt.js";
import Back from "../components/Back.vue";
import Search from "../components/Search.vue";
import SqlQuery from "../components/SqlQuery.vue";

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

function toChildTableData(childTable) {
  console.log("update table")
  tmitt.emit("clickChildTable", childTable)
  store.setChildTable(childTable)
}



</script>

<template>
  <div class="table-list">
    <div class="search">
      <Back class="back"></Back>
      <Search :display-type="3" class="search-input"></Search>
      <SqlQuery class="sql-query"></SqlQuery>
    </div>
    <a-scrollbar outer-style="height:calc(100% - 80px);" style="height:100%;overflow: auto;" class="table-item-list">
      <ChildTableList :child-table="item" v-for="item in childTableList" :key="item"
                      @click="toChildTableData(item)"></ChildTableList>
    </a-scrollbar>
    <!--    <div class="bottom">
          <TablePage :total="200" class="table-page"></TablePage>
        </div>-->
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
  height: calc(100% - 80px);
}

.bottom {
  display: flex;
  height: 35px;
  width: 100%;
  align-items: center;
  justify-content: space-between;
}

.table-page {
  width: 200px

}


</style>