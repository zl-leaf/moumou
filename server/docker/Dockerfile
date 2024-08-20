FROM alibaba-cloud-linux-3-registry.cn-hangzhou.cr.aliyuncs.com/alinux3/golang:1.19.4
LABEL authors="yipzale"

ARG CONFIG_FILE=/var/www/moumou/bootstrap/config/config_prod.yml

WORKDIR /go/src/github.com/moumou/server
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go build -o /var/www/bootstrap/app

WORKDIR /var/www/bootstrap
RUN mkdir  -p build/log
RUN mkdir  -p build/config
COPY $CONFIG_FILE ./config/
COPY ./docker/entrypoint.sh .
RUN chmod a+x ./entrypoint.sh

EXPOSE 8888
ENTRYPOINT ["./entrypoint.sh"]