<template>
  <v-card class="mx-auto" max-width="400">
    <v-img
      src="https://home.ikebukuro.kokosil.net/wp-content/uploads/2019/11/image13052x175_sub3.png"
      height="300px">
    </v-img>
    <MoleculesLoginUserTextArea @organisumValue="userIdData"/>
    <MoleculesLoginPasswordTextArea @organisumValue="passwordData"/>
    <MoleculesLoginButton :clickFunc="login"/>
    <MoleculesSignUpButton :clickFunc="signup"/>
  </v-card>
</template>
<script>
import router from '@/router';
import { mapActions, mapGetters } from 'vuex';

import MoleculesLoginUserTextArea from '@/components/molecules/auth/share/UserTextArea.vue';
import MoleculesLoginPasswordTextArea from '@/components/molecules/auth/share/PasswordTextArea.vue';
import MoleculesLoginButton from '@/components/molecules/auth/login/LoginButton.vue';
import MoleculesSignUpButton from '@/components/molecules/auth/login/SignUpButton.vue';

export default {
  name: 'Login',
  components: {
    MoleculesLoginUserTextArea,
    MoleculesLoginPasswordTextArea,
    MoleculesLoginButton,
    MoleculesSignUpButton,
  },
  data: function() {
    return {
      userId: '',
      password: '',
    };
  },
  methods: {
    ...mapActions({
      testLogin: 'login/testLogin'}),
    ...mapGetters({
      token: 'login/token'}),
    login() {
      const credential = {userId: this.userId, password: this.password};
      this.testLogin(credential);
    },
    signup() {
      router.push('/signup');
    },
    passwordData(value) {
      this.password = value;
    },
    userIdData(value) {
      this.userId = value;
    },
  },
}
</script>
