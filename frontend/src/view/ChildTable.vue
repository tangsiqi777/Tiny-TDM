<script setup>
import {onMounted, ref} from "vue";
import ChildTableList from "../components/ChildTableItem.vue";
import {ListChildTable} from "../../wailsjs/go/main/App.js";
import {Store} from "../store.js";

console.log("ChildTable List\n\n\n\n\n\n")
const store = Store()

let childTableList = ref([])


onMounted(() => {
  displayChildTable()
})

function displayChildTable() {
  let database = store.database
  let superTable = store.superTable
  ListChildTable(store.conn.conn, database, superTable).then((childTables) => {
    childTableList.value = childTables
  })
}

function toChildTableData(table) {
  console.log("update table")
  store.setChildTable(table)
}


</script>

<template>
  <div class="table-list">
    <ChildTableList :child-table="item" v-for="item in childTableList" :key="item"
                    @click="toChildTableData(item)"></ChildTableList>
  </div>
</template>

<style scoped>

</style>