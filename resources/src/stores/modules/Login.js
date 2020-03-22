import getters from '@/stores/getters/Login'
import mutations from '@/stores/mutations/Login.js'
import actions from '@/stores/actions/Login'

export default {
    namespaced: true,
    getters,
    mutations,
    actions,
    state: {
      userId: '',
      password: '',
      token: '',
    },
};
