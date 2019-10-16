<template>
<div class="createduty">
<el-form ref="form" :model="form" label-width="100px">
  <el-form-item label="Request_person">
    <el-input v-model="form.Request_person"></el-input>
  </el-form-item>
  <el-form-item label="Alert_person">
      <el-input v-model="form.Alert_person"></el-input>
  </el-form-item>
   <el-form-item label="Begindate">
    <el-col :span="11">
      <el-date-picker type="date" placeholder="选择日期" v-model="form.Begindate" style="width: 100%;"></el-date-picker>
    </el-col>
  </el-form-item>
   <el-form-item label="Enddate">
    <el-col :span="11">
      <el-date-picker type="date" placeholder="选择日期" v-model="form.Enddate" style="width: 100%;"></el-date-picker>
    </el-col>
  </el-form-item>
  <el-form-item>
    <el-button type="primary" @click="onSubmit">立即创建</el-button>
    </el-form-item>
</el-form>
</div>
</template>



<script>
import * as API from '@/api/dbinstance/';


 export default {
     name: 'createduty',
    data() {
      return {
        form: {
            Request_person:'',
            Alert_person:'',
            Begindate:'',
            Enddate:'',
        },
      }
    },
    methods: {
      onSubmit() {
          console.log(this.form)
         API.createDuty(this.form).then((res) => {
          this.$notify({
            title: '创建值班成功',
            message: `创建值班ID为${res.data.id}`,
          });
        } ).catch((error) => {
         this.$notify({
            title: '创建值班失败',
            message:error,
          });
        });
    },
    },
    components: {
  },
 }
</script>