FROM golang

COPY ./cron.go /usr/src/app/main.go

WORKDIR /usr/src/app

CMD ["go", "run", "main.go"]