<script setup>

import {Store} from "../store.js";

let store = Store()

function back() {
  let displayType = store.displayType
  if (displayType > 0) {
    store.setDisplayType(displayType - 1)
  }
}
</script>

<template>
  <div class="back" @click="back()">
    <icon-arrow-left size="25px"/>
  </div>

</template>

<style scoped>

.back {
  width: 45px;
  height: 45px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.back:hover {
  background-color: #E0EAFC;
}


</style>


