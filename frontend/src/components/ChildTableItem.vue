<script setup>

import {ref} from "vue";
import {SingleMitt} from "../mitt.js";
import {Store} from "../store.js";
import {SqlQuery} from "../../wailsjs/go/service/RestSqlService.js";
import {hasError, parseNestedJsonAndGetData} from "../TdengineRestData.js";
import {Message} from "@arco-design/web-vue";

const props = defineProps(['childTable'])
const store = Store()
let selectedCss = ref("")


SingleMitt.on("clickChildTable", (childTableName) => {
  console.log(childTableName);
  if (childTableName === props.childTable) {
    selectedCss.value = 'background: #dadadb;'
  } else {
    selectedCss.value = ""
  }
});

function toChildTableData(childTable) {
  console.log("display table" + childTable)
  SingleMitt.emit("clickChildTable", childTable)
  store.setChildTable(childTable)
}

function getChildTableInfo() {
  SqlQuery(store.conn.conn, "SHOW CREATE TABLE `" + store.database + "`.`" + props.childTable + "`")
      .then((childTableInfo) => {
        console.log("info:" + JSON.stringify(childTableInfo))
        let errorMsg = hasError(childTableInfo);
        if (errorMsg !== "") {
          Message.error({
            id: 'getSuperTableInfo',
            content: errorMsg,
            duration: 2000
          });
          return
        }
        SingleMitt.emit("displayInfo", {"infoType": 3, "data": parseNestedJsonAndGetData(childTableInfo)})
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
  font-size: 18px;
  line-height: 30px;
  text-align: left;
  color: #4d5869;
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