<script setup>
</script>

<template>
  <div class="content">
    <div class="slide">
      <el-tabs :tab-position="'left'" style="height: 200px" class="demo-tabs">
        <el-tab-pane label="User">User</el-tab-pane>
        <el-tab-pane label="Config">Config</el-tab-pane>
        <el-tab-pane label="Role">Role</el-tab-pane>
        <el-tab-pane label="Task">Task</el-tab-pane>
      </el-tabs>


    </div>
    <div class="main">内容</div>
  </div>

</template>
<style>
.content {
  display: flex;
  height: 100%;
  color: white;
  background: rgba(0, 0, 0, 0.3);
}

.slide {
  width: 5%;
  box-sizing: border-box;

}

.main {
  width: 95%;
  //background-color: #e7d0d0;
  height: 100%;
  background-color: white;
}

.demo-tabs > .el-tabs__content {
  padding: 32px;
  color: #6b778c;
  font-size: 32px;
  font-weight: 600;
}

.el-tabs--right .el-tabs__content,
.el-tabs--left .el-tabs__content {
  height: 100%;
}
</style>




