import Vue from 'vue'
import Router from 'vue-router'

import Signup from '@/components/pages/Signup'
import Login from '@/components/pages/Login'
import Home from '@/components/pages/Home'

Vue.use(Router)

export default new Router({
    routes: [{
            path: '/',
            name: 'Home',
            component: Home
        },
        {
            path: '/signup',
            name: 'Signup',
            component: Signup
        },
        {
            path: '/login',
            name: 'Login',
            component: Login
        }
    ]
})
