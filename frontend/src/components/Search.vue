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
  <div class="search-c">
    <a-input class="search-input-c" v-model="searchValue"
             @press-enter="search"
             placeholder="表名称回车搜索"/>
  </div>
</template>

<style scoped>
.search-c {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.search-input-c {
  width: 85%;
  height: 100%;
}

</style>


