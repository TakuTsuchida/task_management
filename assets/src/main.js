import Vue from 'vue';
import App from './App'
import router from '@/router/router.js';
import vuetify from './plugins/vuetify';

Vue.use(Vuetify);

new Vue({
    el: '#app',
    router,
    vuetify,
    render: h => h(App)
})