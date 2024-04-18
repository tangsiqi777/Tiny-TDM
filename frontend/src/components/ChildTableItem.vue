<script setup>

const props = defineProps(['childTable'])

</script>

<template>
  <div class="child-table-item">
    <div class="child-table"><icon-file size="30px" :strokeWidth="3"/></div>
      <div class="name" >{{ props.childTable }}</div>
    <div class="setting">
      <icon-settings size="22px" :strokeWidth="2"/>
    </div>
  </div>
</template>

<style scoped>

.child-table-item {
  min-height: 50px;
  height: 50px;
  display: flex;
  align-items: center;
  min-width: 300px;
  width: 100%;
}

.child-table-item:hover {
  background: linear-gradient(#E0EAFC, #CFDEF3);
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