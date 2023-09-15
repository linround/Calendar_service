# 基础镜像，基于golang的alpine镜像构建--编译阶段
FROM golang:alpine AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64\
    GOPROXY=https://goproxy.cn

WORKDIR /app
COPY go.sum .
COPY go.mod .

RUN go mod tidy
COPY . .

RUN go build -o /app/main main.go

FROM alpine

COPY --from=builder /app/main /app/main
COPY --from=builder /app/config.yaml /config.yaml





# RUN 是在构建时执行的，CMD命令是在开启时执行的
ENTRYPOINT ["./app/main"]
# 只能在开启服务之后在暴露端口
EXPOSE 8888