FROM golang

COPY ./main.go /usr/src/app/main.go

WORKDIR /usr/src/app

EXPOSE 4000

CMD ["go", "run", "main.go"]