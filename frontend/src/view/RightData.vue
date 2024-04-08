<script setup>

import {PageData1} from "../../wailsjs/go/main/App.js";
import {ref} from "vue";
import router from "../router.js";

let headList = ref([])
let pageDataList = ref([])


let database = router.currentRoute.value.query.database
let table = router.currentRoute.value.query.table
pageData(database, table)

function pageData(database, table) {
  PageData1(database, table).then((pageData) => {
    console.log(JSON.stringify(pageData))
    headList.value = pageData.HeaderList
    pageDataList.value = pageData.Data
    console.log("headList:" + JSON.stringify(headList.value))
    console.log("pageDataList:" + JSON.stringify(pageDataList.value))
  })
}

</script>

<template>
  <el-table :data="pageDataList" border style="width: 100%">
    <el-table-column :prop="item" :label="item" v-for="item in  headList"/>
    <el-table-column fixed="right" label="Operations" width="120">
      <template #default>
        <el-button link type="primary" size="small">Edit</el-button>
      </template>
    </el-table-column>
  </el-table>
</template>

<style scoped>

</style>