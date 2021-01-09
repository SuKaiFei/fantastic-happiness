FROM alpine:latest

LABEL sukaifei kaifeisu@gmail.com

RUN mkdir -p /app
RUN mkdir -p /app/logs

WORKDIR /app

ADD fantastic-happiness /app/fantastic-happiness
COPY configs /app/configs

EXPOSE 8000
EXPOSE 9000

CMD ["./fantastic-happiness", "-conf","configs", "-log.dir","./logs"]