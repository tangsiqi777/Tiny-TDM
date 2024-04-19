<script setup>

import {tmitt} from "../mitt.js";
import {ref} from "vue";

let selectedCss = ref("")
let selected = ref(false)


function displaySqlQuery() {
  if (selected.value === true) {
    selected.value = false
    selectedCss.value = ''
  } else {
    selected.value = true
    selectedCss.value = 'background: #ACE2E1;'

  }

  tmitt.emit("displaySqlQuery", selected.value);
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
  //background-color: #F0EBE3;
  background-color: #ACE2E1;
}

</style>