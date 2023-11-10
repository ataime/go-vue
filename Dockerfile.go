# 使用官方 Golang 镜像
FROM golang:latest

# 设置工作目录
WORKDIR /app

# 将源代码复制到工作目录
COPY src/ ./

# 下载所有依赖项
RUN go mod download


# 构建应用
RUN go build -o main .

# 暴露端口
EXPOSE 9090

# 运行 go 程序
CMD ["./main"]
