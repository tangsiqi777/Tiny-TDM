<script setup>

// import TableItem from "./components/SuperTableItem.vue";
import {onMounted} from 'vue'
import Database from "./view/Database.vue";
import SuperTable from "./view/SuperTable.vue";
import ChildTable from "./view/ChildTable.vue";
import {Store} from "./store.js";
import {storeToRefs} from "pinia";
import RightData from "./view/RightData.vue";
import Blank from "./view/Blank.vue";
import Server from "./view/Server.vue";

const store = Store()
// 0 显示连接， 1显示数据库， 2显示超级表， 3显示子表
let {displayType} = storeToRefs(store)

onMounted(() => {
  dragControllerDiv();
})

function back() {
  console.log("displayType" + displayType.value)
  if (displayType.value > 0) {
    store.setDisplayType(displayType.value - 1)
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
      <div class="left">
        <div class="left-bottom">
          <div class="back">
            <img src="./assets/images/back.png" alt="" v-if="displayType>0" @click="back">
          </div>
          <div class="path">

          </div>
        </div>
        <a-scrollbar class="left-top" style="height: calc(100vh - 6px);overflow-y: auto;">
          <div v-if="displayType===0">
            <Server></Server>
          </div>
          <div v-if="displayType===1">
            <Database></Database>
          </div>
          <div v-if="displayType===2">
            <SuperTable></SuperTable>
          </div>
          <div v-if="displayType===3">
            <ChildTable></ChildTable>
          </div>
        </a-scrollbar>
      </div>


      <div class="resize" title="收缩侧边栏"></div>

      <div class="right-data">
        <div v-if="displayType===3">
          <RightData></RightData>
        </div>
        <div v-if="displayType!==3">
          <Blank></Blank>
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
  height: 100%;
  //background-color: #e7d0d0; height: 100%; background-color: white; display: flex; min-width: 600px; overflow: hidden;
}

.left {
  min-width: 300px;
  width: calc(30% - 6px); /*左侧初始化宽度*/
  height: 100vh;
  background: #FFFFFF;
  float: left;
}

.right-data {
  overflow-y: scroll;
  float: left;
  width: 70%; /*右侧初始化宽度*/
  height: 100%;
  background: #fff;
  //box-shadow: -1px 4px 5px 3px rgba(0, 0, 0, 0.11);
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


.left-top {
  width: 100%;
  min-width: 300px;
}

.left-bottom {
  min-width: 150px;
  width: 100%;
  text-align: left;
  //background-color: #67c23a;
  height: 40px;
  display: flex;
}

.back {
  height: 100%;
  display: flex;
  align-items: center;
  margin-left: 10px;
  //background-color: red;
  width: calc(20% - 10px);
}

.back img {
  height: 60%;
}

.path {
  height: 100%;
  width: 80%;
  //background-color: #79bbff;
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
  //border: #909399 solid 1px;
}

/*拖拽区鼠标悬停样式*/
.resize:hover {
  background-color: #444444;
}

</style>




