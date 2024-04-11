<script setup>

import {onMounted, reactive} from "vue";
import {ListConnection} from "../../wailsjs/go/service/ConnectionStorageService.js";
import {Store} from "../store.js";
import ServerItem from "../components/ServerItem.vue";

const connectionList = reactive({connectionList: []})
onMounted(() => {
  displayConnection()
})


function displayConnection() {
  ListConnection().then((connections) => {
    console.log("conn:" + JSON.stringify(connections))
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
  <div class="connection-list">
    <ServerItem :connectionName="item.name" v-for="item in connectionList.connectionList" :key="item.id"
                @click="toDatabase(item)"></ServerItem>
  </div>
</template>

<style scoped>

</style>