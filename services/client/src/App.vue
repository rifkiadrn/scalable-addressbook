<template>
  <div id="app">
    <Page/>
  </div>
</template>

<script>
import Page from './components/ThePage.vue'
import 'vue-atlas/dist/vue-atlas.css'

export default {
  name: 'app',
  components: {
    Page
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin: 60px;
}
</style>
