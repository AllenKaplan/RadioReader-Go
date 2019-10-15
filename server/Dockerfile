
FROM golang:1.9.0 as builder

WORKDIR /go/src/github.com/AllenKaplan/RadioReader/server

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -o server -a -installsuffix cgo .

FROM alpine:latest

RUN mkdir /app
WORKDIR /app
COPY --from=builder /go/src/github.com/AllenKaplan/RadioReader/server .

CMD ["./main"]
