<script setup>

import {onMounted, ref} from 'vue'
import Database from "./view/Database.vue";
import SuperTable from "./view/SuperTable.vue";
import ChildTable from "./view/ChildTable.vue";
import {Store} from "./store.js";
import {storeToRefs} from "pinia";
import RightData from "./view/RightData.vue";
import Blank from "./components/Blank.vue";
import AddConnection from "./components/AddConnection.vue";
import Server from "./view/Server.vue";
import {Quit, WindowMaximise, WindowMinimise, WindowUnmaximise} from "../wailsjs/runtime/runtime";
import {tmitt} from "./mitt.js";
import SqlQueryPage from "./view/SqlQueryPage.vue";
// 是否非 Mac 平台
const isNotMac = navigator.userAgent.toUpperCase().indexOf('MAC') < 0;
// 是否最大化
const isMaximised = ref(false);

const store = Store()
// 0 显示连接， 1显示数据库， 2显示超级表， 3显示子表
let {displayType} = storeToRefs(store)
let sqlQuery = ref(false)

let searchValue = ref("")

onMounted(() => {
  dragControllerDiv();
})

tmitt.on("displaySqlQuery", (selected) => {
  console.log(selected)
  sqlQuery.value = selected
});


function windowChange() {
  if (isMaximised.value === false) {
    isMaximised.value = true
    WindowMaximise()
  } else {
    isMaximised.value = false
    WindowUnmaximise()
  }
}


function dragControllerDiv() {
  const main = document.getElementsByClassName('main');
  const left = document.getElementsByClassName('left');

  const resize = document.getElementsByClassName('resize');

  const right = document.getElementsByClassName('right');


  for (let i = 0; i < resize.length; i++) {
    // 鼠标按下事件
    resize[i].onmousedown = function (e) {
      //颜色改变提醒
      resize[i].style.background = '#5d5a5a';
      const startX = e.clientX;
      resize[i].left = resize[i].offsetLeft;
      // 鼠标拖动事件
      document.onmousemove = function (e) {
        const endX = e.clientX;
        let moveLen = resize[i].left + (endX - startX); // （endx-startx）=移动的距离。resize[i].left+移动的距离=左边区域最后的宽度
        const maxT = main[i].clientWidth - resize[i].offsetWidth; // 容器宽度 - 左边区域的宽度 = 右边区域的宽度

        if (moveLen < 32) moveLen = 32; // 左边区域的最小宽度为32px
        if (moveLen > maxT - 150) moveLen = maxT - 150; //右边区域最小宽度为150px

        resize[i].style.left = moveLen; // 设置左侧区域的宽度

        for (let j = 0; j < left.length; j++) {
          left[j].style.width = moveLen + 'px';
          right[j].style.width = (main[i].clientWidth - moveLen - 10) + 'px';
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
  <!-- windows 定制化窗口按钮 -->
  <div v-if="isNotMac" class="win-tap" style="--wails-draggable:drag">
    <div class="logo">
      <img class="logo-img" src="./assets/images/tmp/tdengine.ico" alt=""/>
      <span>Tiny TDM</span>
    </div>
    <div class="tap">
      <div @click="WindowMinimise" class="minimise">
        <icon-minus size="20px" :strokeWidth="3"/>
      </div>
      <div @click="windowChange" v-if="isMaximised" class="un-maximise">
        <icon-shrink size="20px" :strokeWidth="3"/>
      </div>
      <div @click="windowChange" v-if="!isMaximised" class="maximise">
        <icon-expand size="20px" :strokeWidth="3"/>
      </div>
      <div @click="Quit" class="quit">
        <icon-close size="20px" :strokeWidth="3"/>
      </div>
    </div>
  </div>
  <div class="content">
    <div class="main">
      <div class="left">
        <Server class="left-fun" v-if="displayType===0"></Server>

        <Database class="left-fun" v-if="displayType===1"></Database>

        <SuperTable class="left-fun" v-if="displayType===2"></SuperTable>

        <ChildTable class="left-fun" v-if="displayType===3"></ChildTable>
      </div>

      <div class="resize" title="收缩侧边栏"></div>

      <div class="right-data">
        <div v-if="sqlQuery === false" class="right-data-inner">
          <RightData v-if="displayType===3"></RightData>
          <AddConnection v-if="displayType===0"></AddConnection>
          <Blank v-if="displayType===1 || displayType === 2"></Blank>
        </div>
        <div v-if="sqlQuery === true" class="right-data-inner">
          <SqlQueryPage></SqlQueryPage>
        </div>
      </div>
    </div>

  </div>

</template>
<style>
/*1296db*/
html,
body {
  height: 100%;
}

.win-tap {
  height: 40px;
  background-color: white;
  width: 100%;
  display: flex;
  justify-content: space-between;
  border-bottom: #E0EAFC solid 1px;
}

.logo {
  height: 35px;
  width: 130px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.logo span {
  font-size: 15px;
  font-weight: bold;
  padding-left: 10px;
}

.logo img {
  height: 32px;
  width: 32px;
}

.tap {
  width: 105px;
  display: flex;
  justify-content: space-around;
  align-items: center;
  height: 35px;
}

.tap > div {
  height: 35px;
  width: 35px;
  display: flex;
  justify-content: space-around;
  align-items: center;
}

.minimise:hover {
  background-color: #dadadb;
}

.un-maximise:hover {
  background-color: #dadadb;
}

.maximise:hover {
  background-color: #dadadb;
}

.quit:hover {
  background-color: #f96c6d;
}


.content {
  display: flex;
  height: calc(100% - 40px);
}


.main {
  width: 100%;
  height: 100%;
  display: flex;
  background-color: white;
  min-width: 600px;
  overflow: hidden;
}

.left {
  min-width: 300px;
  width: calc(20% - 6px); /*左侧初始化宽度*/
  height: 100%;
}

.left-fun {
  height: 100%;
}

.right-data {
  width: 80%; /*右侧初始化宽度*/
  height: 100%;
  background: #fff;
}

.right-data-inner{
  height: 100%;
}


.item-img img {
  width: 40px;
  height: 40px;
}


.back img {
  height: 22px;
  width: 22px;
}


/*拖拽区div样式*/
.resize {
  cursor: col-resize;
  float: left;
  position: relative;
  top: 45%;
  border-radius: 5px;
  width: 6px;
  height: 30px;
  background-color: #858383;
  z-index: 9999;
}

/*拖拽区鼠标悬停样式*/
.resize:hover {
  background-color: #444444;
}

</style>




