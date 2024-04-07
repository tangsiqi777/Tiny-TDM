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
  dragControllerDiv();
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


function dragControllerDiv() {
  const resize = document.getElementsByClassName('resize');
  const left = document.getElementsByClassName('left');
  const mid = document.getElementsByClassName('mid');
  const box = document.getElementsByClassName('box');
  for (let i = 0; i < resize.length; i++) {
    // 鼠标按下事件
    resize[i].onmousedown = function (e) {
      //颜色改变提醒
      resize[i].style.background = '#818181';
      const startX = e.clientX;
      resize[i].left = resize[i].offsetLeft;
      // 鼠标拖动事件
      document.onmousemove = function (e) {
        const endX = e.clientX;
        let moveLen = resize[i].left + (endX - startX); // （endx-startx）=移动的距离。resize[i].left+移动的距离=左边区域最后的宽度
        const maxT = box[i].clientWidth - resize[i].offsetWidth; // 容器宽度 - 左边区域的宽度 = 右边区域的宽度

        if (moveLen < 32) moveLen = 32; // 左边区域的最小宽度为32px
        if (moveLen > maxT - 150) moveLen = maxT - 150; //右边区域最小宽度为150px

        resize[i].style.left = moveLen; // 设置左侧区域的宽度

        for (let j = 0; j < left.length; j++) {
          left[j].style.width = moveLen + 'px';
          mid[j].style.width = (box[i].clientWidth - moveLen - 10) + 'px';
        }
      };
      // 鼠标松开事件
      document.onmouseup = function (evt) {
        //颜色恢复
        resize[i].style.background = '#d6d6d6';
        document.onmousemove = null;
        document.onmouseup = null;
        resize[i].releaseCapture && resize[i].releaseCapture(); //当你不在需要继续获得鼠标消息就要应该调用ReleaseCapture()释放掉
      };
      resize[i].setCapture && resize[i].setCapture(); //该函数在属于当前线程的指定窗口里设置鼠标捕获
      return false;
    };
  }
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
    <div class="main box">
      <div class="server-database-table left">

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
      <div class="resize" title="收缩侧边栏"></div>
      <div class="page-data mid">
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


/* 拖拽相关样式 */
/*包围div样式*/
.box {
  width: 100%;
  height: 100%;
  margin: 1% 0px;
  overflow: hidden;
  box-shadow: -1px 9px 10px 3px rgba(0, 0, 0, 0.11);
}

/*左侧div样式*/
.left {
  width: calc(32% - 10px); /*左侧初始化宽度*/
  height: 100%;
  background: #FFFFFF;
  float: left;
}

/*拖拽区div样式*/
.resize {
  cursor: col-resize;
  float: left;
  position: relative;
  top: 45%;
  background-color: #d6d6d6;
  border-radius: 5px;
  margin-top: -10px;
  width: 10px;
  height: 50px;
  background-size: cover;
  background-position: center;
  /*z-index: 99999;*/
  font-size: 32px;
  color: white;
}

/*拖拽区鼠标悬停样式*/
.resize:hover {
  color: #444444;
}

/*右侧div'样式*/
.mid {
  float: left;
  width: 68%; /*右侧初始化宽度*/
  height: 100%;
  background: #fff;
  box-shadow: -1px 4px 5px 3px rgba(0, 0, 0, 0.11);
}


</style>




