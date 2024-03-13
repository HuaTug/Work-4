FROM golang:alpine

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build
#设置容器中的工作目录为`/build` 所有后续的命令和操作将在`build`目录中执行，这有助于组织和管理构建过程中的文件
# 将代码复制到容器中
COPY . .
#将当前构建上下文中的所有文件复制到容器的当前工作目录 即`build'中
# 将我们的代码编译成二进制可执行文件app
RUN go mod tidy

WORKDIR /build

RUN go build -o app .

# 声明服务端口
EXPOSE 10001

# 启动容器时运行的命令
CMD ["./app"]
