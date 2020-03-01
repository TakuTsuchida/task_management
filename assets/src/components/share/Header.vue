<template>
    <v-app-bar app clipped-left>
        <v-app-bar-nav-icon @click.stop="drawer = !drawer" />
        <v-toolbar-title>Application</v-toolbar-title>
        <v-spacer></v-spacer>
        <v-tooltip bottom>
            <template v-slot:activator="{ on }">
                <v-icon v-on="on">mdi-heart</v-icon>
            </template>
            <span>{{ this.username }}</span>
        </v-tooltip>
    </v-app-bar>
</template>
<script>
import axios from 'axios';
export default {
    data: () => ({
        drawer: null,
        username: null,
    }),
    mounted() {
        this.getUserName();
    },
    methods: {
        getUserName() {
            this.$session.start();
            if (this.$session.has("token")) {
              const headers = {'Authorization': `Bearer ${this.$session.get("token")}`}
              axios.get('http://localhost/api/auth/header', {headers}
              )
                  .then(res => (this.username = res["data"]["username"])
                  )}
        },
    }
}
</script>
