import Vue from 'vue';
import Vuex from 'vuex';
import * as actions from '@stores/actions/index.js';
import * as getters from '@stores/getters/index.js';
import * as states from '@stores/states/index.js';
import * as mutations from '@stores/mutations/index.js';

Vue.use(Vuex);

export default new Vuex.Store({
  actions,
  getters,
  states,
  mutations,
});
