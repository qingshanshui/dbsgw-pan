<template>
  <!--  dialog框需要加的-->
  <n-dialog-provider>
    <!--  message消息需要加的-->
    <n-message-provider>
      <!-- 路由匹配视口-->
      <router-view></router-view>
    </n-message-provider>
  </n-dialog-provider>
</template>