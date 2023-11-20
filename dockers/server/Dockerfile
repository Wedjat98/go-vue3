# 使用 Go 镜像
FROM golang:1.16

# 设置工作目录
WORKDIR /backend

# 复制 Go 模块和依赖项
COPY go.mod /backend/
COPY go.sum /backend/
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN go build -o /server

# 暴露端口
EXPOSE 8080

# 启动服务
CMD ["/server"]
