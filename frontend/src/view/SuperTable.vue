<script setup>

import SuperTableItem from "../components/SuperTableItem.vue";
import {ref} from "vue";
import {ListSuperTable} from "../../wailsjs/go/main/App.js";
import {useRouter} from "vue-router";

let superTableList = ref([])

let router = useRouter()
let database = router.currentRoute.value.query.database

displaySuperTable(database)

function displaySuperTable(database) {
  ListSuperTable(database).then((superTables) => {
    superTableList.value = superTables
    console.log("stable" + JSON.stringify(superTables))
  })
}

function toChildTable(superTable) {
  //父路由编程式传参(一般通过事件触发)
  router.push({
    name: 'childTable',
    query: {
      superTable: superTable,
      database: database
    }
  });
}

</script>


<template>
  <div class="table-list">
    <SuperTableItem :superTable="item" v-for="item in superTableList" :key="item"
                    @click="toChildTable(item)"></SuperTableItem>
  </div>
</template>

<style scoped>

</style>