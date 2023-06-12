<script setup lang="ts">
import {reactive, ref, getCurrentInstance} from "vue";
import {useMessage, UploadCustomRequestOptions, useDialog} from 'naive-ui'
import {useRoute} from "vue-router"
import {chunkFile, mergeFile, verifyFile} from "@/api";
import bus from "@/utils/bus";
import {nanoid} from "nanoid";

const dialog = useDialog()
// 获取上下文
const {proxy} = <any>getCurrentInstance()
let route = useRoute()
const message = useMessage()
let uploadRef = ref()
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
const handleClick = () => {
  uploadRef.value?.submit()
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

  let formData = new FormData()
  formData.append('fileName', file.name)
  formData.append('filePath', route.path)
  let verifyFileRes: any = await verifyFile(formData)
  if (verifyFileRes.data.code != 1000) {
    return dialog.warning({
      title: `检测到有同命名文件【${file.name}】，是否覆盖？`,
      positiveText: '确定',
      negativeText: '取消',
      onPositiveClick: () => {
        uploadRequest(file, onError, onFinish)
      }
    })
  }
  await uploadRequest(file, onError, onFinish)
}

// 上传文件请求处理方法
const uploadRequest = async (file: any, onError: any, onFinish: any) => {
  // 创建切片
  let fileChunks: any = [] // 切片集合数组
  let size = 1024 * 1024 * 2; // 2m 切片大小
  let index = 0 // 切片序号
  let files = file.file // file文件

  for (let i = 0; i < files.size; i += size) {
    fileChunks.push({
      hash: index++,
      chunk: files.slice(i, i + size)
    })
  }

  console.log(files, file)
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
  formDatas.append('filePath', route.path)

  let res: any = await mergeFile(formDatas)
  console.log(res, "res")
  if (res.data.code != 1000) {
    onError()
    return message.success("上传失败")
  }
  onFinish()
  message.success("上传成功")
  bus.emit("reload")
}

// 向外暴露函数
defineExpose({show, hide})
</script>

<template>
  <n-modal v-model:show="state.showModal" :mask-closable="false" preset="dialog" title="文件上传">
    <div class="uploads">

      <div class="uploads-item">
        <n-upload multiple ref="uploadRef" :default-upload="false" :custom-request="customRequest">
          <n-button type="info" size="small">选择文件</n-button>
        </n-upload>
        <n-button class="uploads-item-click" type="info" size="small" @click="handleClick">上传文件</n-button>
      </div>

    </div>
  </n-modal>
</template>

<style scoped>
.uploads {
  margin: 20px;
}

.uploads-item {
  position: relative;
}

.uploads-item-click {
  position: absolute;
  top: 0;
  right: 0;
}

</style>