<script setup lang="ts">
import {reactive, defineExpose,ref} from "vue";
import {useMessage, UploadCustomRequestOptions} from 'naive-ui'
import {useRoute} from "vue-router"
import {upload} from "@/api";
import bus from "@/utils/bus";
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
const customRequest = ({
                         file,
                         data,
                         headers,
                         withCredentials,
                         action,
                         onFinish,
                         onError,
                         onProgress
                       }: UploadCustomRequestOptions) => {
  const formData = new FormData()
  if (data) {
    Object.keys(data).forEach((key) => {
      formData.append(
          key,
          data[key as keyof UploadCustomRequestOptions['data']]
      )
    })
  }
  formData.append("file", file.file as File)
  upload(formData, {type: state.value, url: route.path}).then(res => {
    if (res.data.code === 1000) {
      message.success("上传成功")
      bus.emit("reload")
      onFinish()
    } else {
      message.warning(res.data?.data)
      onError()
    }
  }).catch(err => {
    message.warning("上传失败")
    onError()
  })
}

// 向外暴露函数
defineExpose({show, hide})
</script>

<template>
  <n-modal v-model:show="state.showModal" :mask-closable="false" preset="dialog" title="文件上传">
    <div class="uploads">
      <n-grid x-gap="12" :cols="2">
        <n-gi class="uploads-item">
          <n-select v-model:value="state.value"  :options="state.options" />
        </n-gi>
        <n-gi class="uploads-item">
            <n-upload :custom-request="customRequest" multiple :show-file-list="false">
              <n-button type="info" size="small">上传当前目录</n-button>
            </n-upload>
        </n-gi>
      </n-grid>
    </div>
  </n-modal>
</template>

<style scoped>
.uploads{
  margin: 20px;
}
.uploads-item {
  display: flex;
  align-items: center;
}

</style>