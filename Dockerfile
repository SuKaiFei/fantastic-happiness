FROM alpine:latest
RUN mkdir -p /app
WORKDIR /app

ADD main /app/main

EXPOSE 8000
EXPOSE 9000

CMD ["./main","-conf","-configs","-log.dir","./logs"]