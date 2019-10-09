<template>
  <div class="about">
      <el-table
      :data="tableData"
      style="width: 100%">
      <el-table-column
        prop="date"
        label="日期"
        width="180">
      </el-table-column>
      <el-table-column
        prop="name"
        label="值班dba"
        width="180">
      </el-table-column>
            <el-table-column
        prop="tel"
        label="联系电话"
        width="180">
      </el-table-column>
    </el-table>
  </div>
</template>


<script>
// @ is an alias to /src

export default {
  name: 'about',
}
</script>