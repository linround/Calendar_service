# 基础镜像，基于golang的alpine镜像构建--编译阶段
FROM golang:alpine AS builder

# 配置编译环境
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64\
    GOPROXY=https://goproxy.cn

WORKDIR /ucalendarService
# 将依赖复制到容器中
COPY go.sum .
COPY go.mod .

# 安装依赖包
RUN go mod tidy
# 将所有源文件复制到容器中
COPY . .

# 编译代码并输出到 ucalendarService
RUN go build -o /ucalendarService/main main.go

# 使用 alpine 镜像
FROM alpine

# 复制构建好的可执行文件到镜像中
COPY --from=builder /ucalendarService/main /ucalendarService/main

# 丛编译容器中复制配置文件 到容器中
COPY --from=builder /ucalendarService/config.yaml /config.yaml





# RUN 是在构建时执行的，CMD命令是在开启时执行的
# 设置服务入口
ENTRYPOINT ["./ucalendarService/main"]
# 开启服务之后在暴露端口
EXPOSE 8888