<script setup>

const props = defineProps(['superTable', 'num'])

</script>

<template>
  <div class="super-table-item">
    <div class="super-table"><img class="super-table-img" src="../assets/images/tmp/super-table1.png" alt=""></div>
    <div class="name">{{ props.superTable }}</div>
    <div class="setting">
      <img class="setting-img" src="../assets/images/setting.png" alt="">
    </div>
  </div>
</template>

<style scoped>
.super-table-item {
  min-height: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  min-width: 300px;
  width: 100%;
  margin-top: 30px;
}

.super-table {
  min-width: 60px;
  height: 60px;
  width: 10%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.super-table-img {
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