<script setup>

import {DescTable, SqlQuery} from "../../wailsjs/go/service/RestSqlService.js";
import {hasError, restDataToJsonObj} from "../version/restDataHandle.js";
import {Store} from "../util/store.js";
import {SingleMitt} from "../util/mitt.js";
import {Message} from "@arco-design/web-vue";
import {ref} from "vue";

let selectedCss = ref("")

let store = Store();
const props = defineProps(['superTable', 'num'])

SingleMitt.on("clickSuperTable", (superTable) => {
  // 设置选中颜色
  if (selectedCss.value !== "") {
    selectedCss.value = ""
  }
  if (superTable === props.superTable) {
    selectedCss.value = 'background: #f96c6d;'
  }
});

function toChildTable(superTable) {
  let database = store.database
  DescTable(store.conn.conn, database, superTable).then((tableInfoRet) => {
    console.log("superTableInfo:" + JSON.stringify(tableInfoRet))
    let tableInfo = restDataToJsonObj(tableInfoRet)
    if (tableInfo !== undefined && tableInfo !== null && tableInfo.length > 0) {
      store.setPrimaryId(tableInfo[0].field)
    }
    console.log("删除选中超级表")
    store.removeSelectedSuperTable()
    store.setSuperTable(superTable)
    store.setDisplayType(3)
  })
}

function getSuperTableInfo() {
  SqlQuery(store.conn.conn, "SHOW CREATE STABLE `" + store.database + "`.`" + props.superTable + "`")
      .then((stableInfo) => {
        let errorMsg = hasError(stableInfo);
        if (errorMsg !== "") {
          Message.error({
            id: 'getSuperTableInfo',
            content: errorMsg,
            duration: 2000
          });
          return
        }
        SingleMitt.emit("displayInfo", {"infoType": 2, "data": restDataToJsonObj(stableInfo)})
      })
}

function toSuperData(superTable) {
  let database = store.database
  store.setSelectedSuperTable(superTable)
  DescTable(store.conn.conn, database, superTable).then((tableInfoRet) => {
    console.log("superTableInfo:" + JSON.stringify(tableInfoRet))
    let tableInfo = restDataToJsonObj(tableInfoRet)
    if (tableInfo !== undefined && tableInfo !== null && tableInfo.length > 0) {
      store.setPrimaryId(tableInfo[0].field)
    }
    if (store.sqlQuerySelected === true) {
      Message.warning({
        id: 'toChildTableData',
        content: "请退出Sql查询模式",
        duration: 2000
      });
      return;
    }
    store.setChildTable("")
    store.setSuperTable(superTable)
    SingleMitt.emit("clickSuperTable", superTable)
  })
}

</script>

<template>
  <div class="super-table-item">
    <div class="super-table">
      <icon-nav size="30px" :strokeWidth="3"/>
    </div>
    <div class="name" @click.stop="toChildTable(props.superTable)">{{ props.superTable }}</div>
    <div class="info" @click.stop="toSuperData(props.superTable)" :style="selectedCss">
      <icon-find-replace size="22px" :strokeWidth="2"/>
    </div>
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
  width: calc(100% - 150px);
  min-width: 150px;
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

.info-img {
  width: 22px;
  height: 22px;
}

</style>