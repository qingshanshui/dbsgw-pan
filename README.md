
# dbsgw-pan

> 本项目是一个在线云盘，方便个人文件存储，上传，下载，分享等等一些列个人在线存储项目


###  运行

拉取代码后在项目根目录执行如下命令：

```shell
# web（前端）：node
cd web/ # 进入web目录
npm i # 下载依赖
npm run dev # 运行项目

# server（后端）：go
cd server/ # 进入web目录
go env -w GO111MODULE=on # 建议开启GO111MODULE
go mod download # 下载依赖
go run main.go # 运行项目
```

### 部署

```shell
# web（前端）：node
npm run build #打包 （dist目录直接nginx部署）

# server（后端）：go
$ENV:GOOS="linux" # 设置linux打包环境
go build mian.go # go build
nohup ./fiber-layout -mode prod # 服务器 nohup工具 
```

### 问题/需求
- [x] 面包屑导航深度过长会挤压页面【......】
- [ ] 删除文件功能
- [ ] 上传文件超过2m就会超时
- [ ] 图片分享链接盗用

### 代码贡献

不完善的地方，欢迎大家 Fork 并提交 PR！