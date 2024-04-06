<script setup>

import DatabaseItem from "./components/DatabaseItem.vue";
// import TableItem from "./components/SuperTableItem.vue";
import {ListChildTable, ListDatabases, ListSuperTable, PageData1} from "../wailsjs/go/main/App.js";
import {onMounted, ref} from 'vue'
import {ArrowRight} from '@element-plus/icons-vue';
import SuperTableItem from "./components/SuperTableItem.vue";
import ChildTableList from "./components/ChildTableList.vue";


let databaseList = ref([])
let superTableList = ref([])
let childTableList = ref([])
// 0 显示连接， 1显示数据库， 2显示超级表， 3显示子表
let displayType = ref(1)
let selectDatabase = ref("")
let selectTable = ref("")

let headList = ref([])
let pageDataList = ref([])

onMounted(() => {
  ListDatabases().then((databases) => {
    databaseList.value = databases
    console.log("db" + JSON.stringify(databases))
  });
})

function displaySuperTable(database) {
  displayType.value = 2;
  selectDatabase.value = database
  console.log("db" + database)
  ListSuperTable(database).then((superTables) => {
    superTableList.value = superTables
    console.log("stable" + JSON.stringify(superTables))
  })
}

function displayChildTable(superTable) {
  displayType.value = 3;
  ListChildTable(selectDatabase.value, superTable).then((childTables) => {
    childTableList.value = childTables
    console.log("table" + JSON.stringify(childTables))
  })
}

function pageData(table) {
  selectTable.value = table
  PageData1(selectDatabase.value, selectTable.value).then((pageData) => {
    console.log(JSON.stringify(pageData))
    headList.value = pageData.HeaderList
    pageDataList.value = pageData.Data
    console.log("headList:" + JSON.stringify(headList.value))
    console.log("pageDataList:" + JSON.stringify(pageDataList.value))
  })
}


</script>

<template>
  <div class="content">
    <div class="slide">
      <div class="fun-item">
        <img class="item-img" src="./assets/images/server.png" alt="">
      </div>
      <div class="fun-item">
        <img class="item-img" src="./assets/images/log.png" alt="">
      </div>

    </div>
    <div class="main">
      <div class="server-database-table">

        <!--        <div class="server-database-table-title">-->
        <!--                    <div class="search-input">-->
        <!--                      <el-input v-model="input" style="width: 160px" placeholder="Please input"/>-->
        <!--                    </div>-->
        <!--        </div>-->

        <div class="server-database-table-list">
          <div class="server-list" v-if="displayType===0">

          </div>
          <div class="database-list" v-if="displayType===1">
            <DatabaseItem :databaseName="item" v-for="item in databaseList" :key="item"
                          @click="displaySuperTable(item)"></DatabaseItem>
          </div>
          <div class="table-list" v-if="displayType === 2">
            <SuperTableItem :superTable="item" v-for="item in superTableList" :key="item"
                            @click="displayChildTable(item)"></SuperTableItem>
          </div>

          <div class="table-list" v-if="displayType === 3">
            <ChildTableList :child-table="item" v-for="item in childTableList" :key="item"
                            @click="pageData(item)"></ChildTableList>
          </div>
        </div>
        <!--        <div class="pagination-block">-->
        <!--          <el-pagination layout="prev, pager, next" :total="1000"/>-->
        <!--        </div>-->
        <div class="title">
          <el-breadcrumb :separator-icon="ArrowRight">
            <el-breadcrumb-item :to="{ path: '/' }">1</el-breadcrumb-item>
            <el-breadcrumb-item>a</el-breadcrumb-item>
            <el-breadcrumb-item>b</el-breadcrumb-item>
            <el-breadcrumb-item>c</el-breadcrumb-item>
          </el-breadcrumb>
        </div>
      </div>

      <div class="page-data">
        <div>
          <el-table :data="pageDataList" border style="width: 100%">
            <el-table-column :prop="item" :label="item" v-for="item in  headList"/>
            <el-table-column fixed="right" label="Operations" width="120">
              <template #default>
                <el-button link type="primary" size="small">Edit</el-button>
              </template>
            </el-table-column>
          </el-table>
        </div>

      </div>
    </div>


  </div>

</template>
<style>
//1296db
html,
body {
  height: 100%;
}

.content {
  display: flex;
  height: 100%;
  color: #d2cdcd;
  background: rgba(0, 0, 0, 0.3);
}

.slide {
  width: 5%;
  box-sizing: border-box;
  background-color: #909399;
  height: 100%;
  min-width: 90px;

}

.main {
  width: 95%;
  //background-color: #e7d0d0;
  height: 100%;
  background-color: white;
  display: flex;
  min-width: 600px;
}

.server-database-table {
  width: 20%;
  min-width: 300px;
}

.page-data {
  width: 80%;
  overflow-y: scroll;
}


.fun-item {
  width: 100%;
  height: 90px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.item-img {
  width: 35px;
  height: 35px;
}

.item-img img {
  width: 35px;
  height: 35px;
}

.server-database-table-title {
  width: 100%;
  height: 5%;
  background-color: #67c23a;
  min-width: 300px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.title {
  min-width: 150px;
  width: 50%;
  text-align: left;
}

.search-input {
  min-width: 150px;
  width: 50%;
}

.server-database-table-list {
  width: 100%;
  height: 90%;
  min-width: 300px;
  overflow-y: scroll;
}

.pagination-block {
  width: 100%;
  height: 5%;
  background-color: #67c23a;
}

</style>




