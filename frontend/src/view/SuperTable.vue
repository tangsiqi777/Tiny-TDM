<script setup>

import SuperTableItem from "../components/SuperTableItem.vue";
import {ref} from "vue";
import {ListSuperTable} from "../../wailsjs/go/main/App.js";
import {Store} from "../store.js";

console.log("SuperTable List\n\n\n\n\n\n")

let superTableList = ref([])

const store = Store()


displaySuperTable()

function displaySuperTable() {
  let database = store.database
  ListSuperTable(store.conn.conn, database).then((superTables) => {
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