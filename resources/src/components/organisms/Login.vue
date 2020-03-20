<template>
  <v-container fluid>
    <MoleculesLoginUserTextArea/>
    <MoleculesLoginPasswordTextArea/>
  </v-container>
</template>
<script>
import MoleculesLoginUserTextArea from '@/components/molecules/login/UserTextArea.vue'
import MoleculesLoginPasswordTextArea from '@/components/molecules/login/PasswordTextArea.vue'
export default {
  name: 'Login',
  components: {
    MoleculesLoginUserTextArea,
    MoleculesLoginPasswordTextArea,
  }
}
</script>
