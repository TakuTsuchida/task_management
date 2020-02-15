<template>
    <v-app id="inspire">
        <Sidebar />
        <Header />
        <v-content>
            <v-container class="fill-height" fluid>
                <v-row align="center" justify="center">
                    <v-col class="shrink">
                        <v-tooltip right>
                            <template v-slot:activator="{ on }">
                                <v-btn :href="source" icon large target="_blank" v-on="on">
                                    <v-icon large>mdi-code-tags</v-icon>
                                </v-btn>
                            </template>
                            <span>Source</span>
                        </v-tooltip>
                        <v-tooltip right>
                            <template v-slot:activator="{ on }">
                                <v-btn icon large href="https://codepen.io/johnjleider/pen/bXNzZL" target="_blank" v-on="on">
                                    <v-icon large>mdi-codepen</v-icon>
                                </v-btn>
                            </template>
                            <span>Codepen</span>
                        </v-tooltip>
                    </v-col>
                </v-row>
            </v-container>
        </v-content>
        <Footer />
    </v-app>
</template>
<script>
import Footer from "@/components/share/Footer.vue"
import Header from "@/components/share/Header.vue"
import Sidebar from "@/components/share/Sidebar.vue"
export default {
    name: 'App',
    components: {
        Footer,
        Header,
        Sidebar
    },
    props: {
        source: String,
    },
    created() {
        this.$vuetify.theme.dark = true
    },
}
</script>