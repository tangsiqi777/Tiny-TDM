<script setup>

const props = defineProps(['childTable'])

</script>

<template>
  <div class="child-table-item">
    <div class="child-table"><img class="child-table-img" src="../assets/images/tmp/child-table.png" alt=""></div>

      <div class="name">{{ props.childTable }}</div>



    <div class="setting">
      <img class="setting-img" src="../assets/images/setting.png" alt="">
    </div>
  </div>
</template>

<style scoped>

.child-table-item {
  min-height: 30px;
  height: 30px;
  display: flex;
  align-items: center;
  min-width: 300px;
  width: 100%;
  margin-top: 30px;
}

.child-table {
  min-width: 60px;
  height: 60px;
  width: 10%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.child-table-img {
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