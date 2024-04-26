<script setup>

import {PageData1} from "../../wailsjs/go/service/SqlService.js";
import {reactive, ref, watch} from "vue";
import {Store} from "../store.js";
import {storeToRefs} from "pinia";

console.log("Right Data\n\n\n\n\n\n")

const scrollPercent = {
  x: '100%',
  y: 'calc(100% - 120px)'
};

let headList = ref([])
let pageDataList = ref([])
let total = ref(1)
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
pageData()

function pageData() {
  hideIcon.value = false
  query.primaryId = store.primaryId
  console.log(window.hello("hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh"))
  PageData1(store.conn.conn, database.value, childTable.value, query)
      .then((pageData) => {
        console.log(JSON.stringify(pageData))
        headList.value = convertArrayToObject(pageData.HeaderList)
        pageDataList.value = pageData.Data
        total.value = pageData.Total
        console.log("headList:" + JSON.stringify(headList.value))
        console.log("pageDataList:" + JSON.stringify(pageDataList.value))
        setTimeout(() => {
          hideIcon.value = true
        }, 300);
      })
}

function convertArrayToObject(arr) {
  if (arr === null || arr === undefined || arr.length === 0) {
    return
  }
  query.primaryId = arr[0]
  return arr.map(item => {
    return {
      title: item,
      dataIndex: item,
    };
  });
}

// 示例用法 查询前获取数据库元信息 todo
/*const arr = ["ts", "value"];
const result = convertArrayToObject(arr);
console.log(result);*/

// 可以直接侦听一个 ref
watch(childTable, async (newQuestion, oldQuestion) => {
  pageData()
})


function onSelect(dateString, date) {
  console.log('onSelect', dateString, date);
}

function onChange(dateString, date) {
  console.log('onChange: ', dateString, date);
}

function onOk(dateString, date) {
  console.log('onOk: ', dateString, date);
  query.timeEnd = dateString[1]
  query.timeStart = dateString[0]
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

function clearTime(){
  query.timeStart = "";
  query.timeEnd = "";
  pageData()
}

console.log(getCurrentDatePlusHours(24));
console.log(getCurrentDatePlusWeeks(1));


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
      <div>
        <a-spin :hide-icon="hideIcon" size="large"/>
      </div>
    </div>


  </div>
  <div class="data">
    <a-table :columns="headList" :data="pageDataList" :pagination="false" :scroll="scrollPercent" :scrollbar="true"
             style="height: calc(100%);overflow-y: auto;"/>
  </div>

  <div class="page">
    <a-pagination :total="total" show-page-size @page-size-change="changeSize" @change="changeCurrent"
                  v-model:page-size="query.size"
                  v-model:current="query.current"/>
  </div>

</template>


<style scoped>

.select {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  height: 60px;
  min-width: 700px;
}

.quick-time {
  min-width: 230px;
  width: 230px;
  display: flex;
  justify-content: flex-start;
}

.select-time {
  min-width: 400px;
  width: 400px;
  display: flex;
  justify-content: flex-start;
}

.time-order {
  min-width: 70px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.data {
  height: calc(100% - 120px);
}


.page {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 60px;
}
</style>