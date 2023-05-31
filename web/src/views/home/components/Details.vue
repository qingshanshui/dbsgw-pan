<!--文件详情-->
<script setup lang="ts">
import {reactive} from "vue";
import {FileSelectIcon, saveAs, sizeToStr} from '@/utils/utils'
import {GetFileDownload} from '@/api'
import useClipboard from 'vue-clipboard3'
import {useMessage} from 'naive-ui'

const message = useMessage()

const {toClipboard} = useClipboard()
const copy = async (path: string) => {
  try {
    await toClipboard(`${location.origin}/v1/download?path=${path}`)
    message.success("复制成功")
  } catch (e: any) {
    message.warning(e)
  }
}
const handelDownload = (path: string) => {
  state.loading = true
  GetFileDownload({path}).then(blob => {
    saveAs(blob.data, props.detail.name)
    state.loading = false
    message.success("下载成功")
  })
}
const props = defineProps(['detail'])

let state = reactive({
  loading: false
})
</script>

<template>
  <n-space vertical>
    <n-spin :show="state.loading">
      <div class="detail">
        <div class="detail-icon">
          <n-icon :component="FileSelectIcon(props.detail)" size="50" color="rgb(24, 144, 255)"/>
        </div>
        <div class="detail-name">
          {{ props.detail.name }}
        </div>
        <div class="detail-info">
          {{ sizeToStr(props.detail.size) }} • {{ props.detail.time }}
        </div>
        <div class="detail-operation">
          <n-button type="primary" @click="copy(props.detail.path)">
            复制链接
          </n-button>
          <n-button type="info" @click="handelDownload(props.detail.path)">
            下载
          </n-button>
        </div>
      </div>
    </n-spin>
  </n-space>
</template>

<style scoped>
.detail {
  text-align: center;
  padding: 30px;
}

.detail > div {
  padding: 5px;
}

.detail-name {
  font-size: 18px;
  font-weight: bold;
}

.detail-icon {
  margin-right: 5px
}

.detail-operation > button {
  margin-right: 5px
}
</style>