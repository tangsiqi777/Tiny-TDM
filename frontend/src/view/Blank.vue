<script setup>
import {getCurrentInstance, reactive, ref} from 'vue'


const {ctx: _this} = getCurrentInstance()
const visible = ref(false);
const form = reactive({
  Name: '',
  Addr: '',
  Port: '',
  Username: '',
  Password: '',
});

const formRef = ref()
const handleClick = () => {
  visible.value = true;
};
const handleBeforeOk = (done) => {
  formRef.value.setFields({
    name: {
      status: 'error',
      message: 'async name error'
    },
    ip: {
      status: 'error',
      message: 'valid post'
    }
  })
  done(false)
  console.log(form)
  /*  SaveConnection(form).then(ret => {
      done()
      form.Name = ''
      form.Addr = ''
      form.Port = ''
      form.Username = ''
      form.Password = ''
    })*/


}


const handleCancel = () => {
  visible.value = false;
}

</script>

<template>
  <a-button @click="handleClick">新增连接信息</a-button>
  <a-modal v-model:visible="visible" title="新建连接" @cancel="handleCancel" @before-ok="handleBeforeOk">
    <a-form :model="form" ref="formRef">
      <a-form-item field="name" label="连接名称">
        <a-input v-model="form.Name"/>
      </a-form-item>
      <a-form-item field="ip" label="IP地址">
        <a-input v-model="form.Addr"/>
      </a-form-item>
      <a-form-item field="post" label="端口">
        <a-input-number v-model="form.Port" placeholder="6041" :min="1" :max="100000"/>
      </a-form-item>
      <a-form-item field="post" label="用户名">
        <a-input v-model="form.Username" placeholder="root"/>
      </a-form-item>
      <a-form-item field="post" label="密码">
        <a-input v-model="form.Password" placeholder="taosdata"/>
      </a-form-item>
    </a-form>
  </a-modal>
</template>

<style scoped>

</style>