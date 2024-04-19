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


