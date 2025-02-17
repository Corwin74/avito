FROM golang:1.23 AS builder

WORKDIR ${GOPATH}/avito-shop/
COPY . ${GOPATH}/avito-shop/

RUN GOPROXY=https://goproxy.cn make build

RUN apt-get update && apt-get install -y --no-install-recommends \
        ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

EXPOSE 8080


CMD ["bin/shop", "-conf", "configs/config.yaml"]
