import {Folder, AlertCircleSharp, DocumentText} from '@vicons/ionicons5'

/**
 * 匹配 文件后缀 对应的 图标
 * @param obj
 * @constructor
 */
export function FileSelectIcon(obj: any) {
    if (obj.isDir) return Folder
    let icon = AlertCircleSharp
    switch (Array.isArray(obj.name.match(/.[^.]+$/)) && obj.name.match(/.[^.]+$/)[0]) {
        case ".txt":
            icon = DocumentText
            break;
        default:
            break;
    }
    return icon
}


/**
 * 文件大小  B 到 GB 的转换
 * @param size
 * @constructor
 */
export function sizeToStr(size: number) {
    let sizeStr: string;
    if (size < 0.1 * 1024) { //如果小于0.1KB转化成B
        sizeStr = size.toFixed(2) + "B";
    } else if (size < 0.1 * 1024 * 1024) {//如果小于0.1MB转化成KB
        sizeStr = (size / 1024).toFixed(2) + "KB";
    } else if (size < 0.1 * 1024 * 1024 * 1024) { //如果小于0.1GB转化成MB
        sizeStr = (size / (1024 * 1024)).toFixed(2) + "MB";
    } else { //其他转化成GB
        sizeStr = (size / (1024 * 1024 * 1024)).toFixed(2) + "GB";
    }
    let len = sizeStr.indexOf("\.");
    let dec = sizeStr.substring(len + 1, 2);
    if (dec == "00") {//当小数点后为00时 去掉小数部分
        return sizeStr.substring(0, len) + sizeStr.substring(len + 3, 2);
    }
    return sizeStr;
}

/**
 * 下载文件
 * @param blob 
 * @param filename 
 */
export function saveAs(blob: any, filename: any) {
    let w = window as any
    if (w.navigator.msSaveOrOpenBlob) {
        w.navigator.msSaveBlob(blob, filename)
    } else {
        const link = document.createElement('a')
        const body: any = document.querySelector('body')

        link.href = window.URL.createObjectURL(blob)
        link.download = filename

        // fix Firefox
        link.style.display = 'none'
        body.appendChild(link)

        link.click()
        body.removeChild(link)

        window.URL.revokeObjectURL(link.href)
    }
}
