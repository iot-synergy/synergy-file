FROM nginx:1.25.3-alpine

# Define the project name | 定义项目名称
ARG SERVICE_STYLE=fms
ARG PROJECT_BUILD_SUFFIX=api
# Define the config file name | 定义配置文件名
ARG CONFIG_FILE=fms.yaml


WORKDIR /app
ENV PROJECT=${SERVICE_STYLE}_${PROJECT_BUILD_SUFFIX}
ENV CONFIG_FILE=${CONFIG_FILE}
EXPOSE 80
EXPOSE 9102
COPY ./${PROJECT} ./
COPY ./etc/${CONFIG_FILE} ./etc/
COPY ./deploy/nginx/default.conf /etc/nginx/conf.d/
COPY ./deploy/nginx/entrypoint.sh /docker-entrypoint.d

RUN ["chmod", "+x", "/docker-entrypoint.d/entrypoint.sh"]



