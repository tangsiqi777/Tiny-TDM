<script setup>
const props = defineProps(['connectionName'])


</script>

<template>
  <div class="connection-item">
    <div class="connection"><img class="connection-img" src="../assets/images/tmp/connection3.png" alt="">
    </div>
    <div class="name">
      {{ props.connectionName }}
    </div>
    <div class="setting">
      <img class="setting-img" src="../assets/images/setting.png" alt="">
    </div>
  </div>
</template>

<style scoped>
.connection-item {
  min-height: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  min-width: 300px;
  width: 100%;
  margin-top: 30px;
}

.connection {
  min-width: 60px;
  height: 60px;
  width: 10%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.connection-img {
  width: 30px;
  height: 30px;
}

.name {
  width: 80%;
  min-width: 210px;
  height: 30px;
  font-size: 18px;
  line-height: 30px;
  text-align: left;
  color: #18489d;
}

.setting {
  min-width: 60px;
  width: 10%;
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: flex-start;
}

.setting-img {
  width: 22px;
  height: 22px;
}


</style>