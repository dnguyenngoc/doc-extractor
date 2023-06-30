<template>
    <div class="logo-container">
      <img src="../../assets/images/logo.png" alt="" class="logo"/>
    </div>
  </template>
  
  <script>
  export default {
    name: 'LogoComponent'
  }
  </script>
  
  <style>
    .logo-container {
    position: absolute;
    top: 20px;
    left: 20px;
    }

    .logo {
    width: 250px;
    height: auto;
    }
  </style>