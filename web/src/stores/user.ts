import {defineStore} from "pinia"

/**
 * 登录信息
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
