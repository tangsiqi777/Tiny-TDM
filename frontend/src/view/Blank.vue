<script setup>
import {reactive} from 'vue'
import {SaveConnection} from "../../wailsjs/go/service/ConnectionStorageService.js";

// do not use same name with ref
const form = reactive({
  Id: 0,
  Name: '',
  Addr: '',
  Port: 0,
  Username: '',
  Password: ''
})

const onSubmit = () => {
  SaveConnection(form)
}
</script>

<template>
  <el-form :model="form" label-width="auto" style="max-width: 600px">
    <el-form-item label="id">
      <el-input v-model="form.Id"/>
    </el-form-item>
    <el-form-item label="name">
      <el-input v-model="form.Name"/>
    </el-form-item>
    <el-form-item label="Addr">
      <el-input v-model="form.Addr"/>
    </el-form-item>
    <el-form-item label="Port">
      <el-input v-model="form.Port"/>
    </el-form-item>
    <el-form-item label="Username">
      <el-input v-model="form.Username"/>
    </el-form-item>
    <el-form-item label="Password">
      <el-input v-model="form.Password"/>
    </el-form-item>
  </el-form>
  <el-form-item>
    <el-button type="primary" @click="onSubmit">Create</el-button>
    <el-button>Cancel</el-button>
  </el-form-item>
</template>

<style scoped>

</style>