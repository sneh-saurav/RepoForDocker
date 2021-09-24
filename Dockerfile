FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY *.go ./

RUN go build -o /docker-go-print

EXPOSE 8085

CMD [ "/docker-go-print" ]

