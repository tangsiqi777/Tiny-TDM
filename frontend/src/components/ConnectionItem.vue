<script setup>
import {Store} from "../store.js";
import {DeleteConnection} from "../../wailsjs/go/service/ConnectionStorageService.js";
import {Message} from "@arco-design/web-vue";
import {SingleMitt} from "../mitt.js";

const props = defineProps({connection: Object})

function toDatabase(conn) {
  console.log("conn to db" + JSON.stringify(conn))
  const store = Store()
  store.setConn(conn)
  store.setDisplayType(1)
}

function deleteConnection(conn) {
  console.log("conn" + JSON.stringify(conn))
  DeleteConnection(conn.id).then(result => {
    if (result === "ok") {
      Message.success({
        id: 'deleteConnection',
        content: "删除成功",
        duration: 2000
      });
    } else {
      Message.error({
        id: 'deleteConnection',
        content: result,
        duration: 2000
      });
    }
    SingleMitt.emit("displayConnection", null)
  })
}

function getConnectionInfo() {
  SingleMitt.emit("displayInfo", {"infoType": 0, "data": props.connection})
}

</script>

<template>
  <div class="connection-item">
    <div class="connection">
      <icon-desktop size="30px" :strokeWidth="3"/>
    </div>
    <div class="name" @click="toDatabase(props.connection)">
      {{ props.connection.name }}
    </div>
    <div class="info" @click.stop="getConnectionInfo()">
      <icon-info-circle size="22px" class="info-img" :strokeWidth="2"/>
    </div>
    <div class="setting" @click.stop="deleteConnection(props.connection)">
      <icon-delete size="22px" class="setting-img" :strokeWidth="2"/>
    </div>
  </div>
</template>

<style scoped>
.connection-item {
  min-height: 50px;
  height: 50px;
  display: flex;
  align-items: center;
  min-width: 300px;
  width: 100%;
}

.connection-item:hover {
  background: #dadadb;
}

.connection {
  min-width: 50px;
  height: 50px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.name {
  width: calc(100% - 150px);
  min-width: 150px;
  height: 30px;
  font-size: 18px;
  line-height: 30px;
  text-align: left;
  color: #4d5869;
  overflow: hidden;
}

.setting {
  width: 50px;
  height: 50px;
  min-width: 50px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.setting:hover {
  background: #f96c6d;
}

.setting-img {
  width: 22px;
  height: 22px;
}

.info {
  width: 50px;
  height: 50px;
  min-width: 50px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.info:hover {
  background: #f96c6d;
}

.info-img {
  width: 22px;
  height: 22px;
}


</style>