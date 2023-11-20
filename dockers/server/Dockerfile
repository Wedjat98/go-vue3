# 使用 Go 镜像
FROM golang:latest

# 设置工作目录
WORKDIR /backend

# 复制 Go 模块和依赖项
COPY backend/go.mod /backend/
COPY backend/go.sum /backend/
RUN go mod download

# 复制源代码
COPY backend/ .

# 构建应用
RUN go build -o /server

# 暴露端口
EXPOSE 8080

# 启动服务
CMD ["/server"]
