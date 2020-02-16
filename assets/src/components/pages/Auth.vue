<template>
    <v-container fluid>
        <v-row>
            <v-col cols="12">
                <v-row justify="center">
                    <v-img src="@/images/image_auth.png" aspect-ratio="1" class="grey lighten-2" max-width="500" max-height="330" />
                </v-row>
            </v-col>
            <v-col cols="12">
                <v-form ref="form" v-model="valid" lazy-validation>
                    <v-row justify="center">
                        <v-col cols="6" md="2">
                            <v-text-field v-model="credentials.username" :counter="70" label="USER ID" maxlength="70" required />
                        </v-col>
                    </v-row>
                    <v-row justify="center">
                        <v-col cols="6" md="2">
                            <v-text-field v-model="credentials.password" :counter="70" label="PASSWORD" maxlength="70" required />
                        </v-col>
                    </v-row>
                    <v-row justify="center">
                        <v-btn class="green white--text" :disabled="!valid" @click="login">SIGN UP</v-btn>
                    </v-row>
                </v-form>
            </v-col>
        </v-row>
    </v-container>
</template>
<script>
import axios from 'axios';
import router from '../../router'
export default {
    name: 'Auth',
    data: () => ({
        credentials: {},
        valid: true,
        loading: false
    }),
    methods: {
        login() {
            axios.post('http://localhost/api/auth', this.credentials)
                .then(res => { // responseの略
                    this.$session.start();
                    this.$session.set('token', res.data.token);
                    router.push('/');
                    // eslint-disable-next-line
                })
        }
    }
}
</script>