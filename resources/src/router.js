import Vue from 'vue';
import Router from 'vue-router';

// import Signup from '@/components/pages/Signup'
import Login from '@/components/organisms/Login'
// import Home from '@/components/pages/Home'

Vue.use(Router)

export default new Router({
  routes: [
  {
            path: '/login',
            name: 'Login',
            component: Login
        },
    ]
})
