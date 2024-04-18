<script setup>
import {ref} from "vue";
import {tmitt} from "../mitt.js";
import {Store} from "../store.js";

let searchValue = ref("")
const props = defineProps(['displayType'])
const store = Store()

function search(KeyboardEvent) {
  console.log(JSON.stringify(KeyboardEvent))
  if (props.displayType === 2) {
    store.setSuperTableSearch(searchValue.value)
    tmitt.emit("searchSuperTable", searchValue.value);
    return;
  }
  if (props.displayType === 3) {
    store.setChildTableSearch(searchValue.value)
    tmitt.emit("searchChildTable", searchValue.value);
  }
}
</script>

<template>
  <div class="search">
    <a-input class="search-input" size="large" v-model="props.searchValue"
             @press-enter="search"
             placeholder="输入子表或超级表名称"/>
  </div>
</template>

<style scoped>
.search {
  width: 100%;
  height: 100%;
}
</style>


