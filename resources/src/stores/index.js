import Vue from 'vue';
import Vuex from 'vuex';
import createPersistedState from 'vuex-persistedstate';

import login from '@/stores/modules/Login';

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    login,
  },
  plugins: [createPersistedState({ storage: window.sessionStorage })],
})
