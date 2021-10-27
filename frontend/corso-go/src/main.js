import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
import VueRouter from 'vue-router'
import router from "./router"
import moment from "moment"
import "moment/locale/it"
import VueMoment from "vue-moment"
import store from './store/store'

Vue.config.productionTip = false

moment.locale('it')

Vue.use(VueRouter)
Vue.use(VueMoment,{
  moment
})

new Vue({
  store,
  vuetify,
  router,
  render: h => h(App)
}).$mount('#app')
