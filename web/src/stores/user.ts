import {defineStore} from "pinia"

/**
 * 权限路由表
 * @methods setUserInfos 设置权限路由表
 */
export const userInfos = defineStore('user', {
    state: () => ({
        DirOrDetail: false,// 文件夹/文件
    }),
    actions: {
        // 设置用户数据
        async setDirOrDetail(data: any) {
            this.DirOrDetail = data
        },
    }
})
