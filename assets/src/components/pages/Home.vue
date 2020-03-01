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
//import UserProfile from '@/components/users/UserProfile.vue'
import router from '../../router';
export default {
    name: 'Home',
    components: {
        Sidebar,
        Header,
        Footer,
    },
    data: () => ({
        credentials: {},
        valid: true,
        loading: false
    }),
    created() {
        this.checkLoggedIn();
    },
    methods: {
        checkLoggedIn() {
            this.$session.start();
            if (!this.$session.has("token")) {
                router.push("/login");
            }
        }
    }
}
</script>