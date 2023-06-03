<!--面包屑导航条-->
<script setup lang="ts">
import {Home} from '@vicons/ionicons5'
import {useRoute, useRouter} from "vue-router"
import {watch, reactive, computed, onMounted,ref} from "vue"
import {useMessage} from 'naive-ui'
import Upload from "@/views/home/components/Upload.vue";

import {userInfos} from "@/stores/user";

let users = userInfos()

let route = useRoute()
let routes = useRouter()
const message = useMessage()

// 监听路由变化/变化面包屑导航
watch(route, () => {
  routeToUrl()
})

// upload上传文件的ref
let uploads = ref();

// 数据状态
let state = reactive({
  routeList: [] as any[],
})

// 上传文件
const handelButton = () => {
  uploads.value.show()
}

// 判断上传文件按钮是否显示
let isUpload = computed(() => {
  return !!localStorage.getItem('token') && !users.DirOrDetail
})

// 面包屑导航事件
const handelRoute = (obj: any) => {
  routes.push(obj.path)
}

// 路由转换成面包屑
function routeToUrl() {
  let routers = route.params.path
  let arr = []
  let str = ""
  for (const routersItem of routers) {
    arr.push({href: routersItem, path: str += "/" + routersItem})
  }
  state.routeList = arr
}

// 页面加载完成后
onMounted(() => {
  routeToUrl()
})

</script>

<template>
  <div class="navs">
    <n-breadcrumb class="navs-breadcrumb">
      <n-breadcrumb-item @click="handelRoute({ path: '/' })">
        <n-icon :component="Home"/>
        <span>主页</span>
      </n-breadcrumb-item>
      <n-breadcrumb-item v-for="item in state.routeList" @click="handelRoute(item)">
        {{ item.href }}
      </n-breadcrumb-item>
    </n-breadcrumb>
    <div class="upload" v-if="isUpload">
      <n-button type="info" size="small" @click="handelButton">上传文件</n-button>
    </div>
    <Upload ref="uploads" />
  </div>

</template>

<style scoped>
.navs {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.navs-breadcrumb {
  margin: var(--margin) 0;
  flex: 1;
  overflow: hidden;
}


/deep/ a,
a:hover {
  text-decoration: none;
}
</style>