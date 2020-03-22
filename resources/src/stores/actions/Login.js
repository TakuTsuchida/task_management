import axios from 'axios';
import router from '@/router'

export default {
  testLogin({commit}, credential) {
    // TODO http://localhost/api は共通化するべきである。
    axios.post('http://localhost/api/login', credential)
      .then(res => {
        if (res.status === 200 && 'token' in res.data) {
          const userData = {
            userId: credential.userId,
            password: credential.password,
            token: res.data.token,
          }
          commit('setUser', userData);
          router.push('/')
        }
      })
  },
}
