<script setup>

import {PageData1} from "../../wailsjs/go/main/App.js";
import {reactive, ref, watch} from "vue";
import {Store} from "../store.js";
import {storeToRefs} from "pinia";

console.log("Right Data\n\n\n\n\n\n")

let headList = ref([])
let pageDataList = ref([])
let total = ref(1)
const store = Store()
let {database, childTable} = storeToRefs(store)
let query = reactive({
  timeOrder: 0,
  timeStart: "",
  timeEnd: "",
  primaryId: "",
  size: 50,
  current: 1,
  timeRange: -1,
})
pageData()

function pageData() {
  PageData1(store.conn.conn, database.value, childTable.value, query)
      .then((pageData) => {
        console.log(JSON.stringify(pageData))
        headList.value = convertArrayToObject(pageData.HeaderList)
        pageDataList.value = pageData.Data
        total.value = pageData.Total
        console.log("headList:" + JSON.stringify(headList.value))
        console.log("pageDataList:" + JSON.stringify(pageDataList.value))
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

// 示例用法
const arr = ["ts", "value"];
const result = convertArrayToObject(arr);
console.log(result);

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
    console.log("rrr")
    query.timeOrder = 0
    pageData()
  }
}

function queryDay(mouseEvent) {
  query.timeStart = getCurrentDatePlusHours(24);
  pageData()

}

function queryWeek(mouseEvent) {
  query.timeStart = getCurrentDatePlusWeeks(1);
  pageData()
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
    pageData()
    return
  }
  if (value === 1) {
    query.timeStart = getCurrentDatePlusWeeks(1);
    pageData()
  }
}

console.log(getCurrentDatePlusHours(24));
console.log(getCurrentDatePlusWeeks(1));


</script>
<template>
  <div class="select">
    <div class="quick-time">
      <a-radio-group type="button" @change="changeRange" v-model="query.timeRange">
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
      />
      <div class="time-order">
        <icon-arrow-down size="20px" :strokeWidth="3" v-if="query.timeOrder === 0" @click="changeTimeOrder"/>
        <icon-arrow-up size="20px" :strokeWidth="3" v-if="query.timeOrder === 1" @click="changeCurrent"/>
        <!--        <icon-arrow-up />-->
      </div>
    </div>


  </div>
  <div class="data">
    <a-table :columns="headList" :data="pageDataList" :pagination="false"
             style="height: calc(100vh - 120px);overflow-y: auto;"/>
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

}

.quick-time {
  width: 20%;
  min-width: 180px;
  display: flex;
  justify-content: space-around;
}

.select-time {
  width: 50%;
  min-width: 450px;
  display: flex;
  justify-content: flex-start;
}

.time-order {
  width: 10%;
  min-width: 90px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.data {
  height: calc(100vh - 120px);

}


.page {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 60px;
}
</style>