# 第一阶段：构建 Go 应用
FROM golang:1.23 AS builder

# 设置国内 Go 模块代理
ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /app
COPY . .

# 下载依赖并构建（关闭 CGO）
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -o crm ./

# 第二阶段：运行时镜像 - 精简版
FROM alpine:3.19
WORKDIR /app

# 设置时区环境变量
ENV TZ=Asia/Shanghai
# 安装时区数据并设置时区
RUN apk add --no-cache tzdata && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

# 拷贝执行文件与配置
COPY --from=builder /app/crm /app/crm
COPY config/app.yml /app/config/app.yml

# 端口（与配置一致可改为 8080；此处保留）
EXPOSE 8088

# 启动命令（保留迁移逻辑，依赖 /bin/sh）
CMD ["/bin/sh", "-c", "exec /app/crm"]