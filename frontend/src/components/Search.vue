<script setup>
import {ref} from "vue";
import {SingleMitt} from "../util/mitt.js";
import {Store} from "../util/store.js";

let searchValue = ref("")
const props = defineProps(['displayType'])
const store = Store()

function search(KeyboardEvent) {
  console.log(JSON.stringify(KeyboardEvent))
  if (props.displayType === 2) {
    store.setSuperTableSearch(searchValue.value)
    SingleMitt.emit("searchSuperTable", searchValue.value);
    return;
  }
  if (props.displayType === 3) {
    store.setChildTableSearch(searchValue.value)
    SingleMitt.emit("searchChildTable", searchValue.value);
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
  width: 90%;
  height: 100%;
}

</style>


