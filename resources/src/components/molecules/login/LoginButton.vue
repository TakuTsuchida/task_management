<template>
  <v-row justify="center">
      <v-col class="text-center" cols="6" md="2">
          <Button
            :label="label" :color="color"
            :isSmall="isSmall" :isLarge="isLarge"
            :isDisabled="isRequired" :clickFunc="clickFunc"/>
      </v-col>
  </v-row>
</template>
<script>
import Button from '@/components/atoms/Button.vue';
export default {
  name: 'LoginButtonArea',
  components: {
    Button,
  },
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
