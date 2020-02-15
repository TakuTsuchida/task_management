import 'vuetify/dist/vuetify.min.css'
import 'material-design-icons-iconfont/dist/material-design-icons.css'
import '@mdi/font/css/materialdesignicons.css'

import Vue from 'vue'
import App from './App.vue'
import Vuetify from 'vuetify'
// import Router from 'vue-router'

const vuetifyOptions = {}

Vue.config.productionTip = false
Vue.use(Vuetify, {
    iconfont: 'md'
}) // Vue.user(Router)

new Vue({
    render: h => h(App),
    vuetify: new Vuetify(vuetifyOptions)

}).$mount('#app')