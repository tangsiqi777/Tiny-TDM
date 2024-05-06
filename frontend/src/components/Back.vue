<script setup>

import {Store} from "../util/store.js";
import {SingleMitt} from "../util/mitt.js";

let store = Store()

function back() {
  let displayType = store.displayType
  if (displayType > 0) {
    store.setDisplayType(displayType - 1)
    console.log("click back" + store.displayType)
  }
  if (displayType < 2) {
    SingleMitt.emit("displaySqlQuery", false);
  }
}
</script>

<template>
  <div class="back-c" @click="back()">
    <icon-arrow-left size="28px"/>
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


