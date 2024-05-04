<script setup>
import {reactive, ref} from 'vue'
import {checkNum, checkStrLen,} from '../valid.js'
import {SaveConnection} from "../../wailsjs/go/service/ConnectionStorageService.js";
import {SingleMitt} from "../mitt.js";
import {Message} from "@arco-design/web-vue";

const visible = ref(false);
const form = reactive({
  Name: '',
  Addr: '',
  Port: '',
  Username: '',
  Password: '',
  // 默认3.0
  Version: 'v3.0'
});

const formRef = ref()
const handleClick = () => {
  visible.value = true;
};
const handleBeforeOk = (done) => {
  // 校验字段
  if (!checkStrLen(form.Name, 20)) {
    formRef.value.setFields({
      name: {
        status: 'error',
        message: '名称不为空最长为20个字符'
      }
    })
    done(false)
    return
  }

  if (!checkStrLen(form.Addr, 20)) {
    formRef.value.setFields({
      addr: {
        status: 'error',
        message: '地址不为空最长为20个字符'
      }
    })
    done(false)
    return
  }
  if (!checkNum(form.Port)) {
    formRef.value.setFields({
      port: {
        status: 'error',
        message: '端口号不为空'
      }
    })
    done(false)
    return
  }

  if (!checkStrLen(form.Username, 100)) {
    formRef.value.setFields({
      username: {
        status: 'error',
        message: '用户名不为空'
      }
    })
    done(false)
    return
  }

  if (!checkStrLen(form.Password, 100)) {
    formRef.value.setFields({
      password: {
        status: 'error',
        message: '密码不为空'
      }
    })
    done(false)
    return
  }
  console.log(form)
  SaveConnection(form).then(errMsg => {
    if ('ok' === errMsg) {
      done()
      form.Name = ''
      form.Addr = ''
      form.Port = ''
      form.Username = ''
      form.Password = ''
      Message.success('新增连接成功')
      SingleMitt.emit("displayConnection", null)
      return
    }
    Message.warning('新增连接失败' + errMsg)
    done()
  })
}


const handleCancel = () => {
  visible.value = false;
}

</script>

<template>
  <div class="add-connection" @click="handleClick">
    <div>
      <icon-link :size="80"/>
      <div>新增连接信息</div>
    </div>
  </div>

  <a-modal v-model:visible="visible" title="新建连接" @cancel="handleCancel" @before-ok="handleBeforeOk">
    <a-form :model="form" ref="formRef">
      <a-form-item field="name" label="连接名称">
        <a-input v-model="form.Name"/>
      </a-form-item>
      <a-form-item field="addr" label="IP地址">
        <a-input v-model="form.Addr"/>
      </a-form-item>
      <a-form-item field="post" label="端口">
        <a-input-number v-model="form.Port" placeholder="6041" :min="1" :max="100000"/>
      </a-form-item>
      <a-form-item field="username" label="用户名">
        <a-input v-model="form.Username" placeholder="root"/>
      </a-form-item>
      <a-form-item field="password" label="密码">
        <a-input v-model="form.Password" placeholder="taosdata"/>
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<style scoped>


.add-connection {
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 20px;
  line-height: 50px;
  color: #7d7d7f;
}

</style>