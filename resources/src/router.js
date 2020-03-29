import Vue from 'vue';
import Router from 'vue-router';

// auth
import SignUp from '@/components/organisms/auth/SignUp'
import Login from  '@/components/organisms/auth/Login'
// import Home from '@/components/pages/Home'

Vue.use(Router)

export default new Router({
  routes: [
  {
            path: '/login',
            name: 'Login',
            component: Login
  },
  {
            path: '/signup',
    name: 'SignUp',
            component: SignUp
        },
    ]
})
