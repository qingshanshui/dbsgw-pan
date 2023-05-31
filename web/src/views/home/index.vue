<script setup lang="ts">
import Headers from "./layout/Header.vue"
import Navs from "./layout/Nav.vue"
import Mains from "./layout/Main.vue";
</script>

<template>
  <div class="container">
    <Headers/>
    <Navs/>
    <Mains/>
  </div>
</template>

<style scoped>
.container {
  background-color: var(--bag)
}
</style>