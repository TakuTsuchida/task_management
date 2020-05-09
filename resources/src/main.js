import Vue from 'vue';
import App from './App.vue';
import Vuetify from 'vuetify';
import 'material-design-icons-iconfont/dist/material-design-icons.css'
import 'vuetify/dist/vuetify.min.css';

import router from './router/index';
import store from './stores/index';
import VueSession from 'vue-session';

const vuetifyOptions = {};
Vue.config.productionTip = false;

Vue.use(Vuetify)
new Vuetify({
  icons: {
    iconfont: 'mdi',
  }
});

Vue.use(VueSession)

new Vue({
  router,
  store,
  render: h => h(App),
  vuetify: new Vuetify(vuetifyOptions),
}).$mount('#app')
