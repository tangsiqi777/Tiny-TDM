<script setup>

import DatabaseItem from "../components/DatabaseItem.vue";
import {onMounted, ref} from "vue";
import {useRouter} from "vue-router";
import {Store} from "../util/store.js";
import Back from "../components/Back.vue";
import {ListDatabases} from "../../wailsjs/go/service/RestSqlService.js";
import {hasError, restDataToJsonObj} from "../version/restDataHandle.js";
import {Message} from '@arco-design/web-vue';

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
    console.log("databases:" + JSON.stringify(databases))
    let errorMsg = hasError(databases);
    if (errorMsg !== "") {
      Message.error({
        id: 'displayDatabase',
        content: conn.name + " : " + errorMsg,
        duration: 4000
      });
      return;
    }

    databaseList.value = restDataToJsonObj(databases).filter(database => database.name !== "information_schema" && database.name !== "performance_schema")
  });
}


</script>

<template>
  <div class="database">
    <Back class="back"></Back>
    <a-scrollbar outer-style="height:calc(100% - 45px);" style="height:100%;overflow: auto;" class="database-list">
      <DatabaseItem :databaseName="item.name" v-for="item in databaseList" :key="item"></DatabaseItem>
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