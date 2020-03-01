<template>
    <v-navigation-drawer v-model="drawer" app clipped>
        <v-list dense>
            <v-list-item link>
                <v-list-item-action>
                    <v-icon>mdi-view-dashboard</v-icon>
                </v-list-item-action>
                <v-list-item-content>
                    <v-list-item-title>Dashboard</v-list-item-title>
                </v-list-item-content>
            </v-list-item>
            <v-list-item link>
                <v-list-item-action>
                    <v-icon>mdi-settings</v-icon>
                </v-list-item-action>
                <v-list-item-content @click="logout">
                    <v-list-item-title>Logout</v-list-item-title>
                </v-list-item-content>
            </v-list-item>
        </v-list>
    </v-navigation-drawer>
</template>
<script>
import router from '../../router';
export default {
    data: () => ({
        drawer: null,
    }),
    methods: {
      logout(){
        this.$session.start();
        this.$session.remove("token");
        router.push('/login');
      }}
}
</script>
