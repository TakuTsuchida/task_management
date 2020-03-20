<template>
  <div>
    <v-text-field v-model="value" :label="label" :maxlength="maxLength" :required="isRequired" />
  </div>
</template>
<script>
export default {
  name: 'TextField',
  props: {
        label: {
      type: String,
    },
    maxLength: {
      type: Number,
    },
    isRequired: {
      type: Boolean,
    },
  },
  data() {
    return {
      value: '',
    };
  },
  watch: {
    value: function() {
      this.$emit('moleculeValue', this.value);
    },
  },
}
</script>
