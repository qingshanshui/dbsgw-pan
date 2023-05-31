<!--头部-->
<script setup lang="ts">
import {Search} from '@vicons/ionicons5'
import logo from '@/assets/logo.png'

</script>

<template>
  <header class="header">
    <div class="header-logo">
      <a href="/">
        <img :src="logo" alt="logo"/>
      </a>
    </div>
    <div class="header-search">
      <n-input round placeholder="">
        <template #prefix>
          <n-icon :component="Search"/>
        </template>
      </n-input>
    </div>
  </header>
</template>

<style scoped>
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: var(--margin) 0;
}

.header-logo {
  width: 100px;
  border-radius: 20%;
  overflow: hidden;
}
</style>
