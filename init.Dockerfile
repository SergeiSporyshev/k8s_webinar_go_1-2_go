FROM golang

COPY ./init.go /usr/src/app/main.go

WORKDIR /usr/src/app

CMD ["go", "run", "main.go"]