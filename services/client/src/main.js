import Vue from 'vue'
import App from './App.vue'
import Va from 'vue-atlas'

Vue.use(Va, 'en')
Vue.config.productionTip = false

new Vue({
  render: h => h(App),
}).$mount('#app')
