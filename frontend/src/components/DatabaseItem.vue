<script setup>

const props = defineProps(['databaseName'])


</script>

<template>
  <div class="database-item">
    <div class="database"><icon-storage size="30px" :strokeWidth="3"/>
    </div>
    <div class="name">
      {{ props.databaseName }}
    </div>
    <div class="setting">
      <icon-settings size="22px" :strokeWidth="2"/>
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
.database-item:hover {
  /*background: linear-gradient(#E0EAFC, #CFDEF3);*/
  background: #dadadb;
}

.database {
  min-width: 60px;
  height: 60px;
  width: 10%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.database-img {
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