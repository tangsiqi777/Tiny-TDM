<script setup>

import {ref} from "vue";
import {SingleMitt} from "../util/mitt.js";
import {Store} from "../util/store.js";
import {SqlQuery} from "../../wailsjs/go/service/RestSqlService.js";
import {hasError, restDataToJsonObj} from "../version/restDataHandle.js";
import {Message} from "@arco-design/web-vue";

const props = defineProps(['childTable'])
const store = Store()
let selectedCss = ref("")


SingleMitt.on("clickChildTable", (childTableName) => {
  // 设置选中颜色
  if (selectedCss.value !== "") {
    selectedCss.value = ""
  }
  if (childTableName === props.childTable) {
    selectedCss.value = 'background: #dadadb;'
  }
});

function toChildTableData(childTable) {
  if(store.sqlQuerySelected === true){
    Message.warning({
      id: 'toChildTableData',
      content: "请退出Sql查询模式",
      duration: 2000
    });
    return;
  }
  store.setChildTable(childTable)
  SingleMitt.emit("clickChildTable", childTable)
}

function getChildTableInfo() {
  SqlQuery(store.conn.conn, "SHOW CREATE TABLE `" + store.database + "`.`" + props.childTable + "`")
      .then((childTableInfo) => {
        let errorMsg = hasError(childTableInfo);
        if (errorMsg !== "") {
          Message.error({
            id: 'getSuperTableInfo',
            content: errorMsg,
            duration: 2000
          });
          return
        }
        SingleMitt.emit("displayInfo", {
          "infoType": 3,
          "data": restDataToJsonObj(childTableInfo)
        })
      })
}

</script>

<template>
  <div class="child-table-item" :style="selectedCss">
    <div class="child-table">
      <icon-file size="30px" :strokeWidth="3"/>
    </div>
    <div class="name" @click="toChildTableData(props.childTable)">{{ props.childTable }}</div>
    <div class="info" @click="getChildTableInfo">
      <icon-info-circle size="22px" :strokeWidth="2"/>
    </div>
  </div>
</template>

<style scoped>

.child-table-item {
  min-height: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  min-width: 300px;
  width: 100%;
}

.child-table-item:hover {
  background: #dadadb;
}

.child-table {
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
}


.name {
  width: calc(100% - 80px);
  min-width: 220px;
  height: 30px;
  font-size: 16px;
  line-height: 30px;
  text-align: left;
  color: black;
  overflow: hidden;
}

.info {
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
}


.info:hover {
  background: #f96c6d;
}


</style>