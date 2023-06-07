<script setup lang="ts">
import {reactive, defineExpose, ref,getCurrentInstance} from "vue";
import {useMessage, UploadCustomRequestOptions} from 'naive-ui'
import {useRoute} from "vue-router"
import {chunkFile, mergeFile} from "@/api";
import bus from "@/utils/bus";
import { nanoid } from "nanoid";
// 获取上下文
const { proxy } = <any>getCurrentInstance()
let route = useRoute()
const message = useMessage()
let state = reactive({
  showModal: false, //model显示隐藏
  value: ref(2),
  options: [
    {
      label: 'api上传',
      value: 1
    },
    {
      label: '当前目录',
      value: 2
    },
  ]
})


// 显示
const show = () => {
  state.showModal = true
}

// 隐藏
const hide = () => {
  state.showModal = false
}

// 文件上传
const customRequest = async ({
                               file,
                               data,
                               headers,
                               withCredentials,
                               action,
                               onFinish,
                               onError,
                               onProgress
                             }: UploadCustomRequestOptions) => {

  // 创建切片
  let fileChunks: any = [] // 切片集合数组
  let size = 1024 * 1024*2; // 2m 切片大小
  let index = 0 // 切片序号
  let files = file.file // file文件

  for (let i = 0; i < files.size; i += size) {
    fileChunks.push({
      hash: index++,
      chunk: files.slice(i, i + size)
    })
  }

  // 生成 文件统一的 nanoid
  let id = nanoid()
  for (let i = 0; i < fileChunks.length; i++) {
    let item = fileChunks[i]
    let formData = new FormData()
    formData.append('fileName', file.name)
    formData.append('fileId', id)
    formData.append('fileChunk', item.chunk)
    formData.append('fileIndex', item.hash)
    await chunkFile(formData)
  }

  let formDatas = new FormData()
  formDatas.append('fileId', id)
  formDatas.append('fileName', file.name)
  formDatas.append('fileIndex', fileChunks.length)
  let res: any = await mergeFile(formDatas)
  console.log(res,"res")
  if (res.data.code != 1000) return  message.success("上传失败")
  message.success("上传成功")

}

// 向外暴露函数
defineExpose({show, hide})
</script>

<template>
  <n-modal v-model:show="state.showModal" :mask-closable="false" preset="dialog" title="文件上传">
    <div class="uploads">
      <n-grid x-gap="12" :cols="2">
        <n-gi class="uploads-item">
          <n-select v-model:value="state.value" :options="state.options"/>
        </n-gi>
        <n-gi class="uploads-item">
          <n-upload :custom-request="customRequest" :show-file-list="false">
            <n-button type="info" size="small">上传当前目录</n-button>
          </n-upload>
        </n-gi>
      </n-grid>
    </div>
  </n-modal>
</template>

<style scoped>
.uploads {
  margin: 20px;
}

.uploads-item {
  display: flex;
  align-items: center;
}

</style>