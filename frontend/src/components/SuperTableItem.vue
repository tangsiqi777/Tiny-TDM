<script setup>

const props = defineProps(['superTable', 'num'])

</script>

<template>
  <div class="super-table-item">
    <div class="super-table"><icon-nav size="30px" :strokeWidth="3"/></div>
    <div class="name">{{ props.superTable }}</div>
    <div class="setting">
      <icon-settings size="22px" :strokeWidth="2"/>
    </div>
  </div>
</template>

<style scoped>
.super-table-item {
  min-height: 50px;
  height: 50px;
  display: flex;
  align-items: center;
  min-width: 300px;
  width: 100%;
}

.super-table-item:hover {
  background: linear-gradient(#E0EAFC, #CFDEF3);
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
  color: #4d5869;
  overflow: hidden;
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