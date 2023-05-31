<script setup lang="ts">
import {  onBeforeUnmount } from "vue";
import { GetFileDownload } from "/@/api";
import { saveAs } from "/@/utils/utils";
import bus from "/@/utils/bus"
import { useMessage } from 'naive-ui'
import useClipboard from 'vue-clipboard3'

const message = useMessage()
const props = defineProps(['detail'])
const { toClipboard } = useClipboard()

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
    GetFileDownload({ path }).then(blob => {
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
        <div class=".col-xs-12" style="margin-bottom: 10px">
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
</style>