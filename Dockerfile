FROM alpine:latest

LABEL sukaifei kaifeisu@gmail.com

RUN mkdir -p /app
RUN mkdir -p /app/logs
# 设置时区为上海
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && apk del tzdata

WORKDIR /app

ADD fantastic-happiness /app/fantastic-happiness
COPY configs /app/configs

EXPOSE 8000
EXPOSE 9000

CMD ["./fantastic-happiness", "-conf","configs", "-log.dir","./logs"]