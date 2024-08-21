FROM alibaba-cloud-linux-3-registry.cn-hangzhou.cr.aliyuncs.com/alinux3/golang:1.19.4
LABEL authors="yipzale"

ARG APP_ENV=prod

# 设置环境变量
ENV ENV ${APP_ENV}

# 复制代码与构建bootstrap目录
WORKDIR /go/src/github.com/moumou/server
COPY . .
RUN mkdir  -p /var/www/bootstrap/config

# 构建
RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go build -o /var/www/bootstrap/app \
    && cp config/config_${APP_ENV}.yml /var/www/bootstrap/config/

# 准备启动文件
WORKDIR /var/www/bootstrap
COPY ./docker/entrypoint.sh .
RUN chmod a+x ./entrypoint.sh

EXPOSE 8888
ENTRYPOINT ["./entrypoint.sh"]