import Vue from 'vue';
import App from './App.vue';
import Vuetify from 'vuetify';
import 'vuetify/dist/vuetify.min.css';

import router from './router';
import store from './stores/index';
import VueSession from 'vue-session';

const vuetifyOptions = {};
Vue.config.productionTip = false;

Vue.use(Vuetify)
Vue.use(VueSession)

new Vue({
  router,
  store,
  render: h => h(App),
  vuetify: new Vuetify(vuetifyOptions),
}).$mount('#app')
