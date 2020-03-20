<template>
    <v-app>
      <v-container fluid>
      <v-row class="mb-6">
        <v-col>
        <Header />
        </v-col>
      </v-row>
      <v-row>
        <v-col>
          <Sidebar />
        </v-col>
        <v-col>
          <v-container fluid>
              <img src="@/images/image_home.jpg" />
          </v-container>
        </v-col>
      </v-row>
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
