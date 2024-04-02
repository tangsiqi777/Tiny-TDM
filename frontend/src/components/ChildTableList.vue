<script setup>

const props = defineProps(['childTable'])

</script>

<template>
  <div class="table-item">
    <div class="name-num">
      <div class="table-name">{{ props.childTable }}</div>
    </div>


    <div class="info">
      <img class="info-img" src="../assets/images/info.png" alt="">
    </div>
    <div class="delete">
      <img class="delete-img" src="../assets/images/delete.png" alt="">
    </div>
  </div>
</template>

<style scoped>

.table-item {
  min-height: 60px;
  min-width: 300px;
  width: 100%;
  height: 60px;
  display: flex;
  align-items: center;
}

.name-num {
  padding-left: 10px;
  width: 80%;
  height: 60px;
  display: flex;
}


.table-name {
  height: 60px;
  line-height: 60px;
  text-align: left;

}

.info {
  height: 60px;
  width: 10%;
  min-width: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.info-img {
  width: 30px;
  height: 30px;
}

.delete {
  height: 60px;
  min-width: 60px;
  width: 10%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.delete-img {
  width: 30px;
  height: 30px;
}

</style>