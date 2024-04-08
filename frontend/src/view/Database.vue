<script setup>

import DatabaseItem from "../components/DatabaseItem.vue";
import {onMounted, ref} from "vue";
import {ListDatabases} from "../../wailsjs/go/main/App.js";
import {useRouter} from "vue-router";

let router = useRouter()
let databaseList = ref([])
onMounted(() => {
  displayDatabase()
})


function displayDatabase() {
  ListDatabases().then((databases) => {
    databaseList.value = databases
    console.log("db" + JSON.stringify(databases))
  });
}

function toSuperTable(database) {
  //父路由编程式传参(一般通过事件触发)
  router.push({
    name: 'superTable',
    query: {
      database: database
    }
  });
}

</script>

<template>
  <div class="database-list">
    <DatabaseItem :databaseName="item" v-for="item in databaseList" :key="item"
                  @click="toSuperTable(item)"></DatabaseItem>
  </div>
</template>

<style scoped>

</style>