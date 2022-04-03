FROM golang:1.18

WORKDIR /go-server/

COPY ./server.go .
COPY ./go.mod .

RUN go build -o server .

CMD ["./server"]
