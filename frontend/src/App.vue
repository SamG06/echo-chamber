<template>
  <main>
    <h1>Echo Chamber</h1>
  <EchoTextArea />
  </main>
</template>

<script setup>
  import EchoTextArea from './components/EchoTextArea.vue'
</script>

<style scoped>
main{
  display: flex;
  align-items: center;
  width: 100%;
  flex-flow: column;
}
</style>
