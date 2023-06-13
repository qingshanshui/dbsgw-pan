<!--文件列表-右键操作-->
<script setup lang="ts">
import {onBeforeUnmount} from "vue";
import {GetFileDownload, Del} from "@/api";
import {saveAs} from "@/utils/utils";
import bus from "@/utils/bus"
import {useMessage} from 'naive-ui'
import useClipboard from 'vue-clipboard3'

const message = useMessage()
const props = defineProps(['detail'])
const {toClipboard} = useClipboard()

const deleteFiles = (path: any, name: any) => {
  Del({filePath: path, fileName: name}).then((res) => {
    console.log(res)
    if (res.data.code != 1000) {
      return message.success("删除失败")
    }
    message.success("删除成功")
    bus.emit("reload")
  })
}
// 复制链接
const copy = async (path: string) => {
  try {
    await toClipboard(`${location.origin}/v1/download?path=${path}`)
    message.success('复制成功')
  } catch (e: any) {
    message.warning(e)
  }
}

// 下载文件
const handelDownload = (path: string) => {
  bus.emit('stateLoading', true)
  GetFileDownload({path}).then(blob => {
    saveAs(blob.data, props.detail.name)
    bus.emit('stateLoading', false)
    message.success('下载成功')
  })
}
onBeforeUnmount(() => {
  bus.off('stateLoading')
})
</script>

<template>
  <div class="rightOperation">
    <div class=".col-xs-12">
      <n-button type="primary" @click="deleteFiles(props.detail.path,props.detail.name)">
        删除
      </n-button>
    </div>
    <div class=".col-xs-12">
      <n-button type="primary" @click="copy(props.detail.path)">
        复制链接
      </n-button>
    </div>
    <div class=".col-xs-12">
      <n-button type="info" @click="handelDownload(props.detail.path)">
        下载
      </n-button>
    </div>
  </div>
</template>

<style scoped>
.rightOperation {
  position: fixed;
  font-size: 20px;
  display: none;
  background-color: #ffffff;
  padding: 20px;
  z-index: 999;
  box-shadow: 0 10px 30px -5px rgb(0 0 0 / 30%);
  border-radius: 6px;
}

.rightOperation > div {
  margin-bottom: 10px
}

.rightOperation > div:last-child {
  margin-bottom: 0px
}
</style>