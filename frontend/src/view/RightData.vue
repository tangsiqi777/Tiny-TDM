<script setup>

import {CountData, PageData} from "../../wailsjs/go/service/RestSqlService.js";
import {reactive, ref} from "vue";
import {Store} from "../util/store.js";
import {storeToRefs} from "pinia";
import {getHead, hasError, restDataToJsonObj} from "../version/restDataHandle.js";
import {Message} from "@arco-design/web-vue";
import {SingleMitt} from "../util/mitt.js";

console.log("into data page")

const scrollPercent = reactive({
  x: '100%',
  y: 'calc(100% - 120px)'
})

let headList = ref([])
let pageDataList = ref([])

let total = ref(0)
let hideIcon = ref(true)
const store = Store()
let {database, childTable} = storeToRefs(store)
let query = reactive({
  timeOrder: 0,
  timeStart: "",
  timeEnd: "",
  primaryId: "",
  size: 50,
  current: 1,
  timeRange: 2,
})

function pageData() {
  hideIcon.value = false
  query.primaryId = store.primaryId
  console.log(window.hello("测试将函数绑定到window"))
  console.log("选中childTable:" + store.childTable)
  console.log("选中superTable:" + store.superTable)
  const tableName = (store.childTable && store.childTable !== "") ? store.childTable : store.selectedSuperTable
  console.log("选中Table:" + tableName + "; query:" + JSON.stringify(query))
  if (!tableName) {
    console.log("tableName为空")
    Message.warning({
      id: 'pageData',
      content: "没有选择表！",
      duration: 2000
    });
    hideIcon.value = true
    return;
  }
  PageData(store.conn.conn, database.value, tableName, query)
      .then((pageData) => {
        let errorMsg = hasError(pageData);
        if (errorMsg !== "") {
          Message.error({
            id: 'pageData',
            content: errorMsg,
            duration: 2000
          });
          return;
        }

        CountData(store.conn.conn, database.value, tableName, query).then((count) => {
          total.value = restDataToJsonObj(count)[0].total
          console.log("count" + JSON.stringify(count) + total.value)
        })
        headList.value = (getHead(pageData))
        pageDataList.value = restDataToJsonObj(pageData)
        if (headList.value && headList.value.length > 5) {
          scrollPercent.x = (headList.value.length / 5) * 100 + '%'
        }
        console.log("headList:" + JSON.stringify(headList.value))
        // console.log("pageDataList:" + JSON.stringify(pageDataList.value))
        setTimeout(() => {
          hideIcon.value = true
        }, 300);
      })
}

SingleMitt.on("clickChildTable", (childTable) => {
  //兄弟组件监听点击事件
  pageData()
});

SingleMitt.on("clickSuperTable", (superTable) => {
  pageData()
});

// 时间选择器回调函数
function onSelect(dateString, date) {
  console.log('onSelect', dateString, date);
}

function onChange(dateString, date) {
  console.log('onChange: ', dateString, date);
}

function onOk(dateString, date) {
  console.log('onOk: ', dateString, date);
  query.timeEnd = dateString[1]
  query.pageDatatimeStart = dateString[0]
  pageData()
}

function changeTimeOrder() {
  if (query.timeOrder === 0) {
    query.timeOrder = 1
    pageData()
    return
  }
  if (query.timeOrder === 1) {
    query.timeOrder = 0
    pageData()
  }
}

function changeSize(pageSize) {
  query.size = parseInt(pageSize)
  pageData()
}

function changeCurrent(current) {
  query.current = parseInt(current)
  pageData()
}

function getCurrentDatePlusHours(h) {
  const currentDate = new Date();
  currentDate.setHours(currentDate.getHours() - h);

  const year = currentDate.getFullYear();
  const month = String(currentDate.getMonth() + 1).padStart(2, '0');
  const day = String(currentDate.getDate()).padStart(2, '0');
  const hours = String(currentDate.getHours()).padStart(2, '0');
  const minutes = String(currentDate.getMinutes()).padStart(2, '0');
  const seconds = String(currentDate.getSeconds()).padStart(2, '0');

  return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
}

function getCurrentDatePlusWeeks(weeks) {
  const currentDate = new Date();
  currentDate.setDate(currentDate.getDate() - (weeks * 7));

  const year = currentDate.getFullYear();
  const month = String(currentDate.getMonth() + 1).padStart(2, '0');
  const day = String(currentDate.getDate()).padStart(2, '0');
  const hours = String(currentDate.getHours()).padStart(2, '0');
  const minutes = String(currentDate.getMinutes()).padStart(2, '0');
  const seconds = String(currentDate.getSeconds()).padStart(2, '0');

  return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
}

function changeRange(value) {
  if (value === 0) {
    query.timeStart = getCurrentDatePlusHours(24);
    query.timeEnd = "";
    pageData()
    return
  }
  if (value === 1) {
    query.timeStart = getCurrentDatePlusWeeks(1);
    query.timeEnd = "";
    pageData()
  }
  if (value === 2) {
    query.timeStart = "";
    query.timeEnd = "";
    pageData()
  }
}

function clearTime() {
  query.timeStart = "";
  query.timeEnd = "";
  pageData()
}
</script>
<template>
  <div class="select">
    <div class="quick-time">
      <a-radio-group type="button" @change="changeRange" v-model="query.timeRange">
        <a-radio :value="2">所有</a-radio>
        <a-radio :value="0">近一天</a-radio>
        <a-radio :value="1">近一周</a-radio>
      </a-radio-group>
    </div>
    <div class="select-time">
      <a-range-picker
          style="width: 360px; margin: 0;"
          show-time
          :time-picker-props="{ defaultValue: ['00:00:00', '00:00:00'] }"
          format="YYYY-MM-DD HH:mm"
          @change="onChange"
          @select="onSelect"
          @ok="onOk"
          @clear="clearTime"
      />
      <div class="time-order">
        <icon-arrow-down size="20px" :strokeWidth="3" v-if="query.timeOrder === 0" @click="changeTimeOrder"/>
        <icon-arrow-up size="20px" :strokeWidth="3" v-if="query.timeOrder === 1" @click="changeTimeOrder"/>
      </div>
      <div class="loading">
        <a-spin :hide-icon="hideIcon" :size="20"/>
      </div>
    </div>
  </div>
  <div class="data">
    <a-table :columns="headList" :data="pageDataList" :pagination="false" :scroll="scrollPercent" :scrollbar="true"
             :bordered="{cell:true}"
             column-resizable
             style="height: calc(100%);overflow-y: auto;"/>
  </div>

  <div class="page">
    <a-pagination :total="total" show-page-size @page-size-change="changeSize" @change="changeCurrent"
                  v-model:page-size="query.size" show-total
                  v-model:current="query.current"/>
  </div>

</template>


<style scoped>

.select {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  height: 45px;
  min-width: 700px;
}

.quick-time {
  min-width: 230px;
  width: 230px;
  display: flex;
  justify-content: center;
}

.select-time {
  min-width: 350px;
  width: 350px;
  display: flex;
  justify-content: flex-start;
}

.time-order {
  min-width: 70px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.loading {
  min-width: 50px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.data {
  height: calc(100% - 90px);
}


.page {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 45px;
}
</style>