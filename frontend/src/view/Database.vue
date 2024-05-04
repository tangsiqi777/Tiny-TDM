<script setup>

import DatabaseItem from "../components/DatabaseItem.vue";
import {onMounted, ref} from "vue";
import {useRouter} from "vue-router";
import {Store} from "../store.js";
import Back from "../components/Back.vue";
import {ListDatabases} from "../../wailsjs/go/service/RestSqlService.js";
import {parseNestedJsonAndGetData} from "../TdengineRestData.js";

console.log("Database List\n\n\n\n\n\n")

let router = useRouter()
let databaseList = ref([])
const store = Store()
onMounted(() => {
  console.log("conn:" + JSON.stringify(store.conn))
  displayDatabase(store.conn.conn)
})


function displayDatabase(conn) {
  ListDatabases(conn).then((databases) => {
    databaseList.value = parseNestedJsonAndGetData(databases)
    console.log("db" + JSON.stringify(parseNestedJsonAndGetData(databases)))
  });
}

function toSuperTable(database) {
  store.setDatabase(database)
  store.setDisplayType(2)
}

</script>

<template>
  <div class="database">
    <Back class="back"></Back>
    <a-scrollbar outer-style="height:calc(100% - 45px);" style="height:100%;overflow: auto;" class="database-list">
      <DatabaseItem :databaseName="item.name" v-for="item in databaseList" :key="item"
                    @click="toSuperTable(item.name)"></DatabaseItem>
    </a-scrollbar>


  </div>

</template>

<style scoped>

.database-list {
  width: 100%;
  height: calc(100% - 35px);
}

.back {
  width: 50px;
  height: 45px;
}


</style>