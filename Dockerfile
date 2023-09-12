# 基础镜像，基于golang的alpine镜像构建--编译阶段
FROM golang:alpine
WORKDIR /app
COPY . .
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64\
    GOPROXY=https://goproxy.cn
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0
RUN go build main.go
EXPOSE 8888
# RUN 是在构建时执行的，CMD命令是在开启时执行的
ENTRYPOINT ["./main"]