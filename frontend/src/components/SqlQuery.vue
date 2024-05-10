<script setup>

import {ref} from "vue";
import {Store} from "../util/store.js";

const store = Store()


const initCss = store.sqlQuerySelected === true ? 'background: #e7f1fd;' : '';

let selectedCss = ref(initCss)


function displaySqlQuery() {
  if (store.sqlQuerySelected === true) {
    store.sqlQuerySelected = false
    selectedCss.value = ''
  } else {
    store.sqlQuerySelected = true
    selectedCss.value = 'background: #e7f1fd;'
  }
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