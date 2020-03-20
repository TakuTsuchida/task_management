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
                            <v-text-field v-model="credentials.password" counter="70" label="PASSWORD" maxlength="70" required />
                        </v-col>
                    </v-row>
                    <div class="my-2">
                       <v-row justify="center">
                            <v-btn class="green white--text" :disabled="!valid" @click="login">LOGIN</v-btn>
                       </v-row>
                    </div>
                    <div class="my-2">
                       <v-row justify="center">
                            <v-btn class="orange white--text" :disabled="!valid" @click="signup">SIGNUP</v-btn>
                       </v-row>
                    </div>
                </v-form>
            </v-col>
        </v-row>
    </v-container>
</template>
<script>
import axios from 'axios';
import router from '../../router'
export default {
    name: 'Login',
    data: () => ({
        credentials: {},
        valid: true,
        loading: false
    }),
    methods: {
        login() {
            axios.post('http://localhost/api/login', this.credentials)
                .then(res => {
                    if (res.status == 200 && 'token' in res.data) {
                        this.$session.start();
                        this.$session.set('token', res.data.token);
                        router.push('/');
                    }
                })
        },
        signup() {
            router.push('/signup');
        }
    }
}
</script>
