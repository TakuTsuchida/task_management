<template>
    <v-app>
        <Header />
        <Sidebar />
        <v-container grid-list-md>
            <img src="@/images/image_home.jpg" />
        </v-container>
        <Footer />
    </v-app>
</template>
<script>
import Sidebar from '@/components/share/Sidebar.vue'
import Header from '@/components/share/Header.vue'
import Footer from '@/components/share/Footer.vue'
import axios from 'axios';
export default {
    name: 'Home',
    components: {
        Sidebar,
        Header,
        Footer
    },
    data: () => ({
        credentials: {},
        valid: true,
        loading: false
    }),
    method: {
        login() {

        }
    },
    // createdの中でaxiosを使います。get()の中のURLは、nginx.confで設定してるので、 /api/ になっています。
    created() {
        axios.get('http://localhost/api/')
            .then(response => {
                console.log(response)
            })
    }
}
</script>