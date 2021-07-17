<template>
  <div class="home">
    <img
      @click="getMessage"
      alt="Vue logo"
      src="../assets/logo.png"
      :style="{ height: '400px' }"
    />
    <HelloWorld :msg="message" />
  </div>
</template>