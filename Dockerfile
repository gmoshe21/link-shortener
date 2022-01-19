FROM golang

WORKDIR /go/src/app


COPY ./conn/conn.go /conn/

COPY . .

EXPOSE 5000

CMD ["go", "run", "main.go"]