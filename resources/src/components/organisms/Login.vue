<template>
  <v-container fluid>
    <MoleculesLoginUserTextArea @organisumValue="userIdData"/>
    <MoleculesLoginPasswordTextArea @organisumValue="passwordData"/>
    <MoleculesLoginButton :clickFunc="login"/>
  </v-container>
</template>
<script>
import router from '@/router';
import { mapActions, mapGetters } from 'vuex';

import MoleculesLoginUserTextArea from '@/components/molecules/login/UserTextArea.vue';
import MoleculesLoginPasswordTextArea from '@/components/molecules/login/PasswordTextArea.vue';
import MoleculesLoginButton from '@/components/molecules/login/LoginButton.vue';
export default {
  name: 'Login',
  components: {
    MoleculesLoginUserTextArea,
    MoleculesLoginPasswordTextArea,
    MoleculesLoginButton,
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
