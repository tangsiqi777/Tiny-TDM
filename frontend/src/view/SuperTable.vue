<script setup>

import SuperTableItem from "../components/SuperTableItem.vue";
import {ref} from "vue";
import {ListSuperTable} from "../../wailsjs/go/main/App.js";
import {Store} from "../store.js";
import {tmitt} from "../mitt.js";
import Search from "../components/Search.vue";
import Back from "../components/Back.vue";


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
  console.log("update super table" + superTableSearch)
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
  <div class="super-table">
    <div class="search-input">
      <Search :display-type="2" class="search-input1"></Search>
    </div>

    <a-scrollbar outer-style="height:calc(100% - 90px);" style="height:100%;overflow: auto;" class="table-list">
      <SuperTableItem :superTable="item" v-for="item in superTableList" :key="item"
                      @click="toChildTable(item)"></SuperTableItem>
    </a-scrollbar>
    <Back class="back"></Back>
  </div>
</template>

<style scoped>
.super-table {
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

.table-list{
  height: calc(100% - 90px);
}

.back {
  height: 45px;
}


</style>