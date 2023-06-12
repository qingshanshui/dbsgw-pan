import request from '@/utils/axios';

// 获取文件列表
export function GetList(data: object) {
    return request({
        url: '/v1/list',
        method: 'POST',
        data: data,
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    });
}

// 获取文件信息
export function GetFile(data: object) {
    return request({
        url: '/v1/get',
        method: 'POST',
        data: data,
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    });
}

// 文件下载
export function GetFileDownload(data: object) {
    return request({
        url: '/v1/download',
        method: 'GET',
        params: data,
        responseType: 'blob',
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    });
}

// 登录接口
export function login(data: object) {
    return request({
        url: '/v1/login',
        method: 'POST',
        data: data,
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    });
}

// 文件上传
export function chunkFile(data: object) {
    return request({
        url: `/v1/upload/chunkFile`,
        method: 'POST',
        data: data,
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    });
}

// 合并切片
export function mergeFile(data: object) {
    return request({
        url: '/v1/upload/mergeFile',
        method: 'POST',
        data: data,
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    });
}

// 检验文件是否存在
export function verifyFile(data: object) {
    return request({
        url: '/v1/upload/verifyFile',
        method: 'POST',
        data: data,
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    });
}

// 删除文件
export function Del(data: object) {
    return request({
        url: '/v1/upload/del',
        method: 'POST',
        data: data,
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    });
}

