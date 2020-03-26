<template>
  <v-row justify="center">
      <v-col class="text-center" cols="8" md="6">
          <v-btn
            :color="color"
            :small="isSmall" :large="isLarge"
            :isDisabled="isRequired" @click="clickFunc">
            {{label}}
          </v-btn>
      </v-col>
  </v-row>
</template>
<script>
export default {
  name: 'LoginButtonArea',
  props: {
    clickFunc: {
      type: Function,
    },
  },
  data: function() {
    return {
      label: 'LOGIN',
      color: 'primary',
      isSmall: false,
      isLarge: true,
      isRequired: false,
    };
  },
}
</script>
