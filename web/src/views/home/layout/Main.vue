<!--主体-->
<script setup lang="ts">
import {reactive, onMounted, watch} from "vue"
import NotFound from "../components/NotFound.vue"
import List from "../components/List.vue"
import Details from "../components/Details.vue"
import {GetList, GetFile} from "@/api"
import {useRoute} from 'vue-router'
import bus from "@/utils/bus";

let route = useRoute()

watch(route, () => {
  initStateFile()
})
let state = reactive({
  list: [] as any[],
  show: false,
  DirOrDetail: true, // true是详情  false是文件夹
})

onMounted(() => {
  initStateFile()
  bus.on("stateLoading", (loading: any) => {
    state.show = loading
  })
})
// 获取文件信息
const initStateFile = async () => {
  state.show = true
  let res = await GetFile({path: route.path === "/" ? "" : route.path})
  if (Array.isArray(res?.data?.data)) {
    if (res?.data?.data[0].isDir) {
      await initStateList()
    } else {
      state.DirOrDetail = true
      state.list = res?.data?.data
    }
  } else {
    state.list = []
  }
  state.show = false
}
// 获取文件列表
const initStateList = async () => {
  state.show = true
  let res = await GetList({path: route.path === "/" ? "" : route.path})
  state.show = false
  state.DirOrDetail = false
  state.list = res?.data?.data || []
}
</script>

<template>
  <div class="mains">
    <n-space vertical>
      <n-spin :show="state.show">
        <NotFound v-if="state.list.length === 0"/>
        <template v-else>
          <template v-if="state.DirOrDetail">
            <!--详情-->
            <Details :detail="state.list[0]"/>
          </template>
          <template v-else>
            <!--列表-->
            <List :list="state.list"/>
          </template>
        </template>
      </n-spin>
    </n-space>
  </div>
</template>

<style scoped>
.mains {
  background-color: var(--background-color);
  padding: 10px 20px;
  border-radius: var(--border-radius);
}
</style>
