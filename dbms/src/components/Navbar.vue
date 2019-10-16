<template>
<el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" router>
  <el-menu-item index="/dbinstances">dbinstance</el-menu-item>
    <el-submenu index="/duty">
          <template slot="title">值班表</template>
  <el-menu-item index="/duty-1">当前值班</el-menu-item>
  <el-menu-item index="/duty-2">新增值班</el-menu-item>
  </el-submenu>
  <el-menu-item index="/sqlrequest" >SQL发布</el-menu-item>
</el-menu>
</template>

<script>
export default {
  name: 'Navbar',
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style>
</style>
