<script setup>
console.log("ljdhalfjhalj")
const props = defineProps(['connectionName'])


</script>

<template>
  <div class="database-item">
    <div class="server"><img class="server-img" src="../assets/images/connection.png" alt="">
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
.database-item {
  min-height: 60px;
  height: 60px;
  display: flex;
  align-items: center;
  min-width: 300px;
  width: 100%;
}

.server {
  min-width: 60px;
  height: 60px;
  width: 10%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.server-img {
  width: 30px;
  height: 30px;
}

.name {
  width: 70%;
  min-width: 120px;
  height: 30px;
  font-size: 18px;
  line-height: 30px;
  text-align: left;
}

.setting {
  min-width: 60px;
  width: 10%;
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.setting-img {
  width: 30px;
  height: 30px;
}

.delete {
  min-width: 60px;
  width: 10%;
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.delete-img {
  width: 30px;
  height: 30px;
}

</style>