<template>
  <v-card class="mx-auto" max-width="400">
    <v-img
      src="https://home.ikebukuro.kokosil.net/wp-content/uploads/2019/11/image13052x175_sub3.png"
      height="300px">
    </v-img>
    <MoleculesLoginUserTextArea @organisumValue="userNameData"/>
    <MoleculesLoginPasswordTextArea @organisumValue="passwordData"/>
    <MoleculesSignUpButton :clickFunc="submitSignUp"/>
  </v-card>
</template>
<script>
import { mapActions} from 'vuex';

import MoleculesLoginUserTextArea from '@/components/molecules/auth/share/UserTextArea.vue';
import MoleculesLoginPasswordTextArea from '@/components/molecules/auth/share/PasswordTextArea.vue';
import MoleculesSignUpButton from '@/components/molecules/auth/signup/SignUpButton.vue';

export default {
  name: 'SignUp',
  components: {
    MoleculesLoginUserTextArea,
    MoleculesLoginPasswordTextArea,
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
      signUp: 'auth/signUp'}),
    submitSignUp() {
      const credential = {userName: this.userName, password: this.password};
      this.signUp(credential);
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
