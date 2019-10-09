<template>
  <div id="app">
<Narbar></Narbar>
<router-view></router-view>
  </div>
</template>

<script>
import Narbar from './components/Navbar.vue'

export default {
  name: 'app',
  components: {
    Narbar
  }
}
</script>

<style>

</style>
