FROM golang:1.24 AS builder

# 设置服务名
ARG service_name=checkout

# 设置 Go 环境
WORKDIR /usr/src/gomall
ENV GOPROXY=https://goproxy.cn,direct
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# 复制整个项目
COPY . .

# 构建
RUN mkdir -p /opt/gomall/${service_name} && \
    go build -v -o /opt/gomall/${service_name}/server ./app/${service_name}

# 运行阶段
FROM alpine:latest

# 设置服务名
ARG service_name=checkout
ENV SERVICE_NAME=${service_name}

# 安装运行时依赖
RUN apk --no-cache add ca-certificates tzdata

# 创建目录并复制二进制文件
RUN mkdir -p /opt/gomall/checkout
COPY --from=builder /opt/gomall/checkout/server /opt/gomall/checkout/

# 设置工作目录
WORKDIR /opt/gomall/checkout

# 创建配置挂载点
VOLUME ["/opt/gomall/checkout/conf"]

# 直接指定二进制文件路径
CMD ["/opt/gomall/checkout/server"]