import 'vuetify/dist/vuetify.min.css'
import 'material-design-icons-iconfont/dist/material-design-icons.css'
import '@mdi/font/css/materialdesignicons.css'

import Vue from 'vue'
import App from './App.vue'
import Vuetify from 'vuetify'
import router from './router'
import VueSession from 'vue-session'

const vuetifyOptions = {}
Vue.config.productionTip = false

Vue.use(Vuetify)
Vue.use(VueSession)

new Vue({
    router,
    render: h => h(App),
    vuetify: new Vuetify(vuetifyOptions)

}).$mount('#app')