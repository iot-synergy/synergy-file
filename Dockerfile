FROM golang:1.21-git  as builder

ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
# ENV GOPRIVATE=github.com/iot-synergy
COPY . /app
WORKDIR /app
RUN go mod tidy --compat=1.21 && go build -o /app/fms-api fms.go


FROM nginx:1.25.3-alpine

# Define the project name | 定义项目名称
ARG PROJECT=fms
# Define the config file name | 定义配置文件名
ARG CONFIG_FILE=fms.yaml

LABEL org.opencontainers.image.authors=${AUTHOR}

WORKDIR /app
ENV PROJECT=${PROJECT}
ENV CONFIG_FILE=${CONFIG_FILE}

COPY --from=builder /app/fms-api ./
COPY --from=builder /app/etc/${CONFIG_FILE} ./etc/
COPY --from=builder /app/deploy/nginx/default.conf /etc/nginx/conf.d/
COPY --from=builder /app/deploy/nginx/entrypoint.sh /docker-entrypoint.d

RUN ["chmod", "+x", "/docker-entrypoint.d/entrypoint.sh"]

EXPOSE 80
EXPOSE 9102

