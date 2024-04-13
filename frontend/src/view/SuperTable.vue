<script setup>

import SuperTableItem from "../components/SuperTableItem.vue";
import {ref} from "vue";
import {ListSuperTable} from "../../wailsjs/go/main/App.js";
import {Store} from "../store.js";
import {tmitt} from "../mitt.js";


console.log("SuperTable List\n\n\n\n\n\n")

let superTableList = ref([])

const store = Store()


tmitt.on("searchSuperTable", (data) => {
  displaySuperTable()
});

displaySuperTable()

function displaySuperTable() {
  let database = store.database
  let superTableSearch = store.superTableSearch
  console.log("update super table"+ superTableSearch)
  ListSuperTable(store.conn.conn, database, superTableSearch).then((superTables) => {
    superTableList.value = superTables
    console.log("stable" + JSON.stringify(superTables))
  })
}

function toChildTable(superTable) {
  store.setSuperTable(superTable)
  store.setDisplayType(3)
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