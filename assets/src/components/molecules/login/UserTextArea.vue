<template>
  <v-row justify="center">
      <v-col cols="6" md="2">
          <TextField @moleculeValue="userIdData" :label="label" :maxLength="maxLength" :isRequired="isRequired"/>
          {{this.userId}}
      </v-col>
  </v-row>
</template>
<script>
// TODO change name loginTextArea.vue
import TextField from '@/components/atoms/TextField.vue';
export default {
  name: 'UserTextArea',
  components: {
    TextField,
  },
  data: function() {
    return {
      label: 'User Id',
      maxLength: 70,
      isRequired: true,
      userId: '',
    };
  },
  methods: {
    userIdData(value) {
      this.userId = value;
    },
  },
}
</script>
