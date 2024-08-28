FROM alibaba-cloud-linux-3-registry.cn-hangzhou.cr.aliyuncs.com/alinux3/golang:1.19.4
LABEL authors="yipzale"

ARG APP_ENV=prod

# 设置环境变量
ENV ENV ${APP_ENV}

# 复制代码
WORKDIR /go/src/github.com/moumou/server
COPY . .

# 构建
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && mkdir -p /var/www/bootstrap \
    && mkdir /var/www/bootstrap/config \
    && cp config/config_${APP_ENV}.yaml /var/www/bootstrap/config/ \
    && go build -o /var/www/bootstrap ./...

# 准备启动文件
WORKDIR /var/www/bootstrap

EXPOSE 8888
ENTRYPOINT ["./server"]