<script setup>
import {Store} from "../util/store.js";
import {SingleMitt} from "../util/mitt.js";
import {SqlQuery} from "../../wailsjs/go/service/RestSqlService.js";
import {hasError, restDataToJsonObj} from "../version/restDataHandle.js";
import {Message} from "@arco-design/web-vue";

const store = Store()
const props = defineProps(['databaseName'])

function toSuperTable(database) {
  store.setDatabase(database)
  store.setDisplayType(2)
}

function getDatabaseInfo() {
  SqlQuery(store.conn.conn, "SHOW CREATE DATABASE `" + props.databaseName + "`").then((databaseInfo) => {
    let errorMsg = hasError(databaseInfo);
    if (errorMsg !== "") {
      Message.error({
        id: 'getDatabaseInfo',
        content: errorMsg,
        duration: 2000
      });
      return
    }
    SingleMitt.emit("displayInfo", {"infoType": 1, "data": restDataToJsonObj(databaseInfo)})
  })
}

</script>

<template>
  <div class="database-item">
    <div class="database">
      <icon-storage size="30px" :strokeWidth="3"/>
    </div>
    <div class="name" @click.stop="toSuperTable(props.databaseName)">
      {{ props.databaseName }}
    </div>
    <div class="info" @click.stop="getDatabaseInfo">
      <icon-info-circle size="22px" :strokeWidth="2"/>
    </div>
  </div>
</template>

<style scoped>
.database-item {
  min-height: 50px;
  height: 50px;
  display: flex;
  align-items: center;
  min-width: 300px;
  width: 100%;
}

.database-item:hover {
  background: #dadadb;
}

.database {
  width: 50px;
  height: 50px;
  display: flex;
  align-items: center;
  justify-content: center;
}


.name {
  width: calc(100% - 100px);
  min-width: 200px;
  height: 30px;
  font-size: 16px;
  line-height: 30px;
  text-align: left;
  color: black;
  overflow: hidden;
}

.info {
  width: 50px;
  height: 50px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.info:hover {
  background: #f96c6d;
}

</style>