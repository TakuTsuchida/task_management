<template>
  <v-btn
    color="color" small="isSmall" large="isLarge",
    disabeld="isDisabled" @click="clickFunc">
    {{ label }}
  </v-btn>
</template>
<script>
export default {
  name: 'Button',
  props: {
    label: {
      type: String,
    },
    isSmall: {
      type: Boolean,
      required: true,
    },
    isLarge: {
      type: Boolean,
      required: true,
    },
    isDisabled: {
      type: Boolean,
      required: true,
    },
    clickFunc: {
      type: Function,
      required: true,
    },
  },
}
</script>
