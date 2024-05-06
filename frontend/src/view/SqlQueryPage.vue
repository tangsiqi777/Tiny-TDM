<script setup>

import {ref} from "vue";
import {Store} from "../util/store.js";
import {isNotEmpty} from "../util/valid.js";
import {SqlQuery} from "../../wailsjs/go/service/RestSqlService.js";
import {getHead, hasError, restDataToJsonObj} from "../version/restDataHandle.js";
import {Message} from "@arco-design/web-vue";

const store = Store()
let headList = ref([])
let pageDataList = ref([])

let selectedDatabase = store.database
let selectedTable = isNotEmpty(store.childTable) ? store.childTable : (isNotEmpty(store.superTable) ? store.superTable : "table1")
let defaultSql = "SELECT * FROM `" + selectedDatabase + "`.`" + selectedTable + "` LIMIT 100"

let sql = ref(defaultSql)


const scrollPercent = {
  x: '100%',
  y: 'calc(100% - 230px)'
};

function subSql() {
  SqlQuery(store.conn.conn, sql.value).then((pageData) => {
    let errorMsg = hasError(pageData);
    if (errorMsg !== "") {
      Message.error({
        id: 'subSql',
        content: errorMsg,
        duration: 2000
      });
      return;
    }
    console.log(JSON.stringify(pageData))
    headList.value = (getHead(pageData))
    pageDataList.value = restDataToJsonObj(pageData)
    console.log("headList:" + JSON.stringify(headList.value))
    console.log("pageDataList:" + JSON.stringify(pageDataList.value))
  })
}

function convertArrayToObject(arr) {
  if (arr === null || arr === undefined || arr.length === 0) {
    return
  }
  return arr.map(item => {
    return {
      title: item,
      dataIndex: item,
    };
  });
}
</script>

<template>
  <div class="sql-query-page-c">
    <div class="info-c">
      <a-alert>1. SQL需要带上数据库限定<br/>
        2. 对于有大写数据库名数据库表的SQL需要对该字段使用``转义<br/>
        3. 为了避免卡顿返回最多显示2000条<br/>
      </a-alert>
    </div>
    <div class="sql-input-c">
      <a-textarea placeholder="输入SQL"
                  v-model="sql" style="height: 150px"
                  allow-clear/>
    </div>
    <div class="sql-sub-c">
      <a-button type="primary" @click="subSql">提交SQL</a-button>
    </div>


    <div class="data-c">
      <a-table :columns="headList" :data="pageDataList" :pagination="false" :scroll="scrollPercent" :scrollbar="true"
               style="height: 100%;overflow-y: auto;"/>
    </div>
  </div>

</template>

<style scoped>


.sql-query-page-c {
  padding-top: 7px;
  margin: 0 auto;
  width: 98%;
  height: 100%;
}

.info-c {
  padding-bottom: 20px;
}

.sql-sub-c {
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: flex-end;

}

.data-c {
  height: calc(100% - 230px);
}


</style>