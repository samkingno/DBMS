<template>
<el-menu :default-active="activeIndex" class="el-menu-demo" mode="horizontal" router>
  <el-menu-item index="/">dbinstance</el-menu-item>
  <el-menu-item index="/about" >值班表</el-menu-item>
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
