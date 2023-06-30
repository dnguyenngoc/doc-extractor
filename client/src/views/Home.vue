<template>
  <div class="hello">
    <div style="padding-bottom: 100px;">
    <LogoComponent />
    </div>
    <div>
      <h1>Hello World</h1>
      <ButtonComponent nameButton="Create Account" buttonClass="button" @click="handleClick" />
      <ButtonComponent nameButton="Login" buttonClass="button-small" @click="handleClick" />
      <ButtonComponent nameButton="Sigup" buttonClass="button-normal" @click="handleClick" />
      <ButtonComponent nameButton="Sigup White" buttonClass="button bg-white" @click="handleClick" />
      <ButtonComponent nameButton="Danger" buttonClass="button-small bg-danger" @click="handleClick" />
      <ButtonComponent nameButton="Warning" buttonClass="button-small bg-warning" @click="handleClick" />
      <ButtonComponent nameButton="Success" buttonClass="button-small bg-success" @click="handleClick" />
      </div>
  </div>
</template>

<script>
  import ButtonComponent from '@/components/button/Button.vue';
  import LogoComponent from '@/components/logo/Logo.vue'
export default {
  name: 'HomeView',
  components: {
    ButtonComponent,
    LogoComponent,
  },
  methods: {
    handleClick() {
    },
  },
}
</script>

<style>
.hello {
  background-color: black;
  /* Additional styling for the .hello class */
}
</style>

