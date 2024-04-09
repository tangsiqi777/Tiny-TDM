<script setup>

import DatabaseItem from "../components/DatabaseItem.vue";
import {onMounted, ref} from "vue";
import {ListDatabases} from "../../wailsjs/go/main/App.js";
import {useRouter} from "vue-router";
import {Store} from "../store.js";

console.log("Database List\n\n\n\n\n\n")

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
  const store = Store()
  store.setDatabase(database)
  store.setDisplayType(2)
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