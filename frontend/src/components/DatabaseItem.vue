<script setup>

const props = defineProps(['databaseName'])


</script>

<template>
  <div class="database-item">
    <div class="server"><img class="server-img" src="../assets/images/database.png" alt="">
    </div>
    <div class="name">
      {{ props.databaseName }}
    </div>
    <div class="setting">
      <img class="setting-img" src="../assets/images/setting.png" alt="">
    </div>
    <div class="blank">

    </div>
  </div>
</template>

<style scoped>
.database-item {
  margin-top: 8px;
  min-height: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  margin-bottom: 15px;
  min-width: 300px;
  width: 100%;
}

.server {
  min-width: 30px;
  height: 30px;
  width: 15%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.server-img {
  width: 30px;
  height: 30px;
}

.name {
  width: 65%;
  min-width: 120px;
  height: 20px;
  font-size: 18px;
  line-height: 30px;
  text-align: left;
}

.setting {
  min-width: 30px;
  width: 15%;
  height: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.setting-img {
  width: 30px;
  height: 30px;
}

.blank{
  width: 5%;
  height: 30px;
}




</style>