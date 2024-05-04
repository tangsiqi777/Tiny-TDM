<script setup>

import {onMounted, reactive} from "vue";
import {ListConnection} from "../../wailsjs/go/service/ConnectionStorageService.js";
import {Store} from "../store.js";
import ServerItem from "../components/ConnectionItem.vue";
import {SingleMitt} from "../mitt.js";

const connectionList = reactive({connectionList: []})
onMounted(() => {
  displayConnection()
})

SingleMitt.on("displayConnection", () => {
  console.log("ljdlajkl")
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


function toDatabase(conn) {
  console.log(JSON.stringify(conn))
  const store = Store()
  store.setConn(conn)
  store.setDisplayType(1)
}

</script>

<template>
  <a-scrollbar outer-style="height:100%;" style="height:100%;overflow: auto;" class="connection-list">
    <ServerItem :connectionName="item.name" v-for="item in connectionList.connectionList" :key="item.id"
                @click="toDatabase(item)"></ServerItem>
  </a-scrollbar>
</template>

<style scoped>
.connection-list {
  width: 100%;
  height: 100%;
}


</style>