<template>
  <v-row justify="center">
      <v-col cols="6" md="2">
          <TextField @moleculeValue="passwordData" :label="label" :maxLength="maxLength" :isRequired="isRequired"/>
          {{this.password}}
      </v-col>
  </v-row>
</template>
<script>
import TextField from '@/components/atoms/TextField.vue';
export default {
  name: 'PasswordTextArea',
  components: {
    TextField,
  },
  data: function() {
    return {
      label: 'Password',
      maxLength: 70,
      isRequired: true,
      password: '',
    };
  },
  methods: {
    passwordData(value) {
      this.password = value;
    },
  },
}
</script>
