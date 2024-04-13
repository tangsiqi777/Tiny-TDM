<script setup>

import {PageData1} from "../../wailsjs/go/main/App.js";
import {ref, watch} from "vue";
import {Store} from "../store.js";
import {storeToRefs} from "pinia";

console.log("Right Data\n\n\n\n\n\n")

let headList = ref([])
let pageDataList = ref([])
const store = Store()
let {database, childTable} = storeToRefs(store)
pageData()

function pageData() {
  PageData1(store.conn.conn, database.value, childTable.value).then((pageData) => {
    console.log(JSON.stringify(pageData))
    headList.value = convertArrayToObject(pageData.HeaderList)
    pageDataList.value = pageData.Data


    console.log("headList:" + JSON.stringify(headList.value))
    console.log("pageDataList:" + JSON.stringify(pageDataList.value))
  })
}

function convertArrayToObject(arr) {
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
}


</script>
<template>
  <div class="select">
    <div class="quick-time">
      <a-button type="primary" style="width: 100px" status="success">24小时</a-button>
      <a-button type="primary" style="width: 100px"  status="normal">近一周</a-button>
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
    </div>


  </div>
  <div class="data">
    <a-table :columns="headList" :data="pageDataList" :pagination="false"
             style="height: calc(100vh - 120px);overflow-y: auto;"/>
  </div>

  <div class="page">
    <a-pagination :total="200" show-page-size/>
  </div>

</template>


<style scoped>

.select {
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 60px;

}

.quick-time{
  width: 30%;
  min-width: 220px;
  display: flex;
  justify-content: space-around;
}
.select-time{
  width: 70%;
  min-width: 600px;
  display: flex;
  justify-content: flex-start;
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