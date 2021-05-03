# 阶段一构建
FROM golang:alpine AS build1

# 设置环境变量
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# 移动到工作目录：/build
WORKDIR /build

# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download

# 将代码复制到容器中
COPY . .

# 将代码编译成二进制可执行文件
RUN go build -o integration-app -tags=jsoniter .


# 创建一个小镜像
FROM scratch

# 建立目录并跳转
WORKDIR /integration-app
WORKDIR /integration-app/image
WORKDIR /integration-app

# 移动配置文件到容器中
COPY ./config.json .

# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=build1 /build/integration-app .

# 暴露端口8086和3306
EXPOSE 8086
EXPOSE 3306

# 运行服务
ENTRYPOINT ["./integration-app"]