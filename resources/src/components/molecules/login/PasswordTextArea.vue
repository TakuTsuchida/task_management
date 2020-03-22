<template>
  <v-row justify="center">
    <v-col cols="6" md="2">
      <v-text-field v-model="value" 
        :label="label" :maxlength="maxLength" :required="isRequired" />
    </v-col>
  </v-row>
</template>
<script>
export default {
  name: 'PasswordTextArea',
  data: function() {
    return {
      label: 'Password',
      maxLength: 70,
      isRequired: true,
      value: '',
    };
  },
  watch: {
    value: function() {
      this.$emit('organisumValue', this.value);
    },
  },
}
</script>
