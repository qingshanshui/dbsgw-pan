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
    <a href="https://beian.miit.gov.cn" target="_blank"
       style="color: #999;position: fixed;left: 50%;bottom: 10px;transition: inherit;transform: translateX(-50%);">鄂ICP备2022012826号</a>
  </div>
</template>

<style scoped>
.container {
  background-color: var(--bag)
}
</style>