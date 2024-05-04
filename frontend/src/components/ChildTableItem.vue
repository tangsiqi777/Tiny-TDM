<script setup>

import {ref} from "vue";
import {SingleMitt} from "../mitt.js";

const props = defineProps(['childTable'])

let selectedCss = ref("")


SingleMitt.on("clickChildTable", (childTableName) => {
  console.log(childTableName);
  if (childTableName === props.childTable) {
    selectedCss.value = 'background: #dadadb;'
  } else {
    selectedCss.value = ""
  }
});

</script>

<template>
  <div class="child-table-item" :style="selectedCss">
    <div class="child-table">
      <icon-file size="30px" :strokeWidth="3"/>
    </div>
    <div class="name">{{ props.childTable }}</div>
    <div class="setting">
      <icon-settings size="22px" :strokeWidth="2"/>
    </div>
  </div>
</template>

<style scoped>

.child-table-item {
  min-height: 60px;
  height: 60px;
  display: flex;
  align-items: center;
  min-width: 300px;
  width: 100%;
}

.child-table-item:hover {
  background: #dadadb;
}

.child-table {
  width: 60px;
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
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


</style>