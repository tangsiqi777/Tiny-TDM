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
  PageData1(database, childTable).then((pageData) => {
    console.log(JSON.stringify(pageData))
    headList.value = pageData.HeaderList
    pageDataList.value = pageData.Data
    console.log("headList:" + JSON.stringify(headList.value))
    console.log("pageDataList:" + JSON.stringify(pageDataList.value))
  })
}

// 可以直接侦听一个 ref
watch(question, async (newQuestion, oldQuestion) => {
  if (newQuestion.includes('?')) {
    loading.value = true
    answer.value = 'Thinking...'
    try {
      const res = await fetch('https://yesno.wtf/api')
      answer.value = (await res.json()).answer
    } catch (error) {
      answer.value = 'Error! Could not reach the API. ' + error
    } finally {
      loading.value = false
    }
  }
})

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