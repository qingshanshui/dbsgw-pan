<script setup lang="ts">
import { Home } from '@vicons/ionicons5'
import { useRoute, useRouter } from "vue-router"
import { watch, reactive, computed, onMounted } from "vue"
import { useMessage, UploadCustomRequestOptions } from 'naive-ui'
import { upload } from "/@/api/index";
let route = useRoute()
let routes = useRouter()
const message = useMessage()
watch(route, () => {
    routeToUrl()
})

let state = reactive({
    routeList: [] as any[],
    showModal: false,
})

const onNegativeClick = () => {
    state.showModal = false
}

const onPositiveClick = () => {
    state.showModal = false
}

const handelButton = () => {
    state.showModal = true
}

let isUpload = computed(() => {
    return localStorage.getItem('token') ? true : false
})


const handelRoute = (obj: any) => {
    routes.push(obj.path)
}

// 路由转换成面包屑
function routeToUrl() {
    let routers = route.params.path
    let arr = []
    let str = ""
    for (const routersItem of routers) {
        arr.push({ href: routersItem, path: str += "/" + routersItem })
    }
    state.routeList = arr
}

onMounted(() => {
    routeToUrl()
})
// api上传
const customRequestApi = ({
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
    upload(formData, { type: 1 }).then(res => {
        if (res.data.code === 1000) {
            message.success("上传成功")
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
// 当前文件上传
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
    upload(formData, { type: 2, url: route.path }).then(res => {
        if (res.data.code === 1000) {
            message.success("上传成功")
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

</script>

<template>
    <div style="display: flex;justify-content: space-between;align-items: center;">
        <n-breadcrumb class="navs">
            <n-breadcrumb-item @click="handelRoute({ path: '/' })">
                <n-icon :component="Home" />
                主页
            </n-breadcrumb-item>
            <n-breadcrumb-item v-for="item in state.routeList" @click="handelRoute(item)">
                {{ item.href }}
            </n-breadcrumb-item>
        </n-breadcrumb>
        <div class="upload" v-if="isUpload">

            <n-button type="info" size="small" @click="handelButton">上传文件</n-button>
        </div>
    </div>
    <n-modal v-model:show="state.showModal" :mask-closable="false" preset="dialog" title="文件上传">
        <div>
            <div style="margin: 10px;">
                <n-upload :custom-request="customRequestApi" multiple :show-file-list="false">
                    <n-button type="info" size="small">api上传</n-button>
                </n-upload>
            </div>
            <div style="margin: 10px;">
                <n-upload :custom-request="customRequest" multiple :show-file-list="false">
                    <n-button type="info" size="small">上传当前目录</n-button>
                </n-upload>
            </div>
        </div>
    </n-modal>
</template>

<style scoped>
.navs {
    margin: var(--margin) 0;
}

/deep/ a,
a:hover {
    text-decoration: none;
}
</style>
