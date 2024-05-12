<script setup>

import {Store} from "../util/store.js";

let store = Store()

function back() {
  const lastDisplayType = store.displayType
  if (lastDisplayType > 0) {
    const newDisplayType = lastDisplayType - 1
    store.setDisplayType(newDisplayType)
    if (lastDisplayType === 3 && newDisplayType === 2) {
      console.log("子表列表返回到超级表列表")
      store.removeChildTable();
    }
  }
  if (store.displayType < 2) {
    store.removeSqlQuerySelected();
  }
}
</script>

<template>
  <div class="back-c" @click="back()">
    <icon-left size="28px"/>
  </div>

</template>

<style scoped>

.back-c {
  display: flex;
  align-items: center;
  justify-content: center;
}

.back-c:hover {
  background-color: #dadadb;
}


</style>


