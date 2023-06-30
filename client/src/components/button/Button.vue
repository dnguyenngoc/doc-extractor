<template>
    <button :class="buttonClass" @click="$emit('click')">
      {{ nameButton }}
    </button>
</template>
  
<script>
  export default {
    name: 'ButtonComponent',
    props: {
      nameButton: {
        type: String,
        default: "Ok",
      },
      buttonClass: {
        type: String,
        default: "button",
      },
    },
  };
</script>
  
<style lang="scss">
    @import "@/assets/styles/base/button.scss";
</style>
