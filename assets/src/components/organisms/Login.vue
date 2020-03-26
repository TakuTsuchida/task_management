<template>
  <v-card class="mx-auto" max-width=344>
    <v-card-actions>
      <MoleculesLoginUserTextArea/>
      <MoleculesLoginPasswordTextArea/>
    </v-card-actions>
  </v-card>
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
