<template>
  <v-card class="mx-auto" max-width="400">
    <v-img
      src="https://home.ikebukuro.kokosil.net/wp-content/uploads/2019/11/image13052x175_sub3.png"
      height="300px">
    </v-img>
    <MoleculesLoginUserTextArea @organisumValue="userNameData"/>
    <MoleculesLoginPasswordTextArea @organisumValue="passwordData"/>
    <MoleculesLoginButton :clickFunc="submitLogin"/>
    <MoleculesSignUpButton :clickFunc="signup"/>
  </v-card>
</template>
<script>
import router from '@/router';
import { mapActions} from 'vuex';

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
      userName: '',
      password: '',
    };
  },
  methods: {
    ...mapActions({
      login: 'auth/login'}),
    submitLogin() {
      const credential = {userName: this.userName, password: this.password};
      this.login(credential);
    },
    signup() {
      router.push('/signup');
    },
    passwordData(value) {
      this.password = value;
    },
    userNameData(value) {
      this.userName = value;
    },
  },
}
</script>
