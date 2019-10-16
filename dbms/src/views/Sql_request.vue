<template>
 <div class="sqlrequest">
<el-form ref="form" :model="form" label-width="100px">
  <el-form-item label="hostname">
    <el-input v-model="form.hostname"></el-input>
  </el-form-item>
  <el-form-item label="port">
      <el-input v-model="form.port"></el-input>
  </el-form-item>
    <el-form-item label="database">
      <el-input v-model="form.database"></el-input>
  </el-form-item>
    <el-form-item label="pg-request">
      <el-input type="textarea"  v-model="form.request"></el-input>
      </el-form-item>
  <el-form-item>
    <el-button type="primary" @click="onSubmit">开始执行</el-button>
    </el-form-item>
</el-form>
</div>
</template>



<script>

import * as API from '@/api/dbinstance/';


  export default {
      name:'sqlrequest',
    data() {
      return {
        form: {
          hostname: '',
          database: '',
          request: '',
        }
      }
    },
  methods: {
      onSubmit() {
          console.log(this.form)
         API.postrequest(this.form).then((res) => {
    if (res.status > 0) {
          this.$notify.error({
            title: '执行失败',
            message: res.error,
          });
        } else{
          this.$notify({
            title: 'sql发布成功',
            message: `sql为${res.data.Request}`,
          });
        }
        } ).catch((error) => {
         this.$notify({
            title: '发布sql失败',
            message:error,
          });
        });
    },
    },
    components: {
  },
  }
</script>