<script setup>

import {DescTable, SqlQuery} from "../../wailsjs/go/service/RestSqlService.js";
import {hasError, parseNestedJsonAndGetData} from "../TdengineRestData.js";
import {Store} from "../store.js";
import {SingleMitt} from "../mitt.js";
import {Message} from "@arco-design/web-vue";

let store = Store();

const props = defineProps(['superTable', 'num'])

function toChildTable(superTable) {
  let database = store.database
  DescTable(store.conn.conn, database, superTable).then((tableInfoRet) => {
    let tableInfo = parseNestedJsonAndGetData(tableInfoRet)
    console.log("tableInfo:" + JSON.stringify(tableInfo))
    if (tableInfo !== undefined && tableInfo !== null && tableInfo.length > 0) {
      console.log("tableInfo:" + tableInfo[0].field)
      store.setPrimaryId(tableInfo[0].field)
    }
    store.setSuperTable(superTable)
    store.setDisplayType(3)
  })
}

function getSuperTableInfo() {
  SqlQuery(store.conn.conn, "SHOW CREATE STABLE `" + store.database + "`.`" + props.superTable + "`")
      .then((stableInfo) => {
        console.log("info:" + JSON.stringify(stableInfo))
        let errorMsg = hasError(stableInfo);
        if (errorMsg !== "") {
          Message.error({
            id: 'getSuperTableInfo',
            content: errorMsg,
            duration: 2000
          });
          return
        }
        SingleMitt.emit("displayInfo", {"infoType": 2, "data": parseNestedJsonAndGetData(stableInfo)})
      })
}

</script>

<template>
  <div class="super-table-item">
    <div class="super-table">
      <icon-nav size="30px" :strokeWidth="3"/>
    </div>
    <div class="name" @click.stop="toChildTable(props.superTable)">{{ props.superTable }}</div>
    <div class="info" @click.stop="getSuperTableInfo">
      <icon-info-circle size="22px" :strokeWidth="2"/>
    </div>
  </div>
</template>

<style scoped>
.super-table-item {
  min-height: 50px;
  height: 50px;
  display: flex;
  align-items: center;
  min-width: 300px;
  width: 100%;
}

.super-table-item:hover {
  background: #dadadb;

}

.super-table {
  min-width: 50px;
  height: 50px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.super-table-img {
  width: 30px;
  height: 30px;
}

.name {
  width: calc(100% - 100px);
  min-width: 200px;
  height: 30px;
  font-size: 18px;
  line-height: 30px;
  text-align: left;
  color: #4d5869;
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

.info-img {
  width: 22px;
  height: 22px;
}

</style>