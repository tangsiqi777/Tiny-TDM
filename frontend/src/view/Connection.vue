<script setup>

import {onMounted, reactive} from "vue";
import {ListConnection} from "../../wailsjs/go/service/ConnectionStorageService.js";
import ServerItem from "../components/ConnectionItem.vue";
import {SingleMitt} from "../mitt.js";

const connectionList = reactive({connectionList: []})
onMounted(() => {
  displayConnection()
})

SingleMitt.on("displayConnection", () => {
  displayConnection()
});


function displayConnection() {
  ListConnection().then((connections) => {
    console.log("conn:" + JSON.stringify(connections))
    connectionList.connectionList.length = 0
    for (let i = 0; i < connections.length; i++) {
      connectionList.connectionList.push(connections[i])
    }
  });
}

</script>

<template>
  <a-scrollbar outer-style="height:100%;" style="height:100%;overflow: auto;" class="connection-list">
    <ServerItem :connection="item" v-for="item in connectionList.connectionList" :key="item.id"></ServerItem>
  </a-scrollbar>
</template>

<style scoped>
.connection-list {
  width: 100%;
  height: 100%;
}


</style>