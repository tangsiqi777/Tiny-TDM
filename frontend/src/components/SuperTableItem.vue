<script setup>

const props = defineProps(['superTable', 'num'])

</script>

<template>
  <div class="super-table-item">
    <div class="super-table">
      <icon-nav size="30px" :strokeWidth="3"/>
    </div>
    <div class="name">{{ props.superTable }}</div>
    <div class="setting">
      <icon-settings size="22px" :strokeWidth="2"/>
    </div>
  </div>
</template>

<style scoped>
.super-table-item {
  min-height: 60px;
  height: 60px;
  display: flex;
  align-items: center;
  min-width: 300px;
  width: 100%;
}

.super-table-item:hover {
  background: #dadadb;

}

.super-table {
  min-width: 60px;
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.super-table-img {
  width: 30px;
  height: 30px;
}

.name {
  width: calc(100% - 120px);
  min-width: 180px;
  height: 30px;
  font-size: 18px;
  line-height: 30px;
  text-align: left;
  color: #4d5869;
  overflow: hidden;
}

.setting {
  width: 60px;
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.setting-img {
  width: 22px;
  height: 22px;
}

</style>