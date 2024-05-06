<script setup>

import {SingleMitt} from "../util/mitt.js";
import {ref} from "vue";

let selectedCss = ref("")
let selected = ref(false)


function displaySqlQuery() {
  if (selected.value === true) {
    selected.value = false
    selectedCss.value = ''
  } else {
    selected.value = true
    selectedCss.value = 'background: #e7f1fd;'

  }

  SingleMitt.emit("displaySqlQuery", selected.value);
}

</script>

<template>
  <div class="sql-query-c" :style="selectedCss" @click="displaySqlQuery">
    <icon-code-square size="28px"/>
  </div>
</template>

<style scoped>
.sql-query-c {
  display: flex;
  align-items: center;
  justify-content: center;
}

.sql-query-c:hover {
  background-color: #e7f1fd;
}

</style>