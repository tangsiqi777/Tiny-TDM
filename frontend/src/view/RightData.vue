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
  PageData1(database.value, childTable.value).then((pageData) => {
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

</script>

<template>
  <a-table :columns="headList" :data="pageDataList"/>
</template>

<style scoped>

</style>