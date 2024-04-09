<script setup>
import {onMounted, ref} from "vue";
import ChildTableList from "../components/ChildTableList.vue";
import {ListChildTable} from "../../wailsjs/go/main/App.js";
import router from "../router.js";

console.log("ChildTable List\n\n\n\n\n\n")

let childTableList = ref([])
let database = ref("")
let superTable = ref("")


onMounted(() => {
  database.value = router.currentRoute.value.query.database
  superTable.value = router.currentRoute.value.query.superTable
  displayChildTable(database.value, superTable.value)
})

function displayChildTable(database, superTable) {
  ListChildTable(database, superTable).then((childTables) => {
    childTableList.value = childTables
  })
}

function toChildTableData(table) {

  //父路由编程式传参(一般通过事件触发)
  router.push({
    name: 'childTableData',
    query: {
      table: table,
      database: database.value
    }
  });
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