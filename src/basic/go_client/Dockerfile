FROM golang:1.13.3

WORKDIR /go/src/app

COPY . .

ENV GO111MODULE on

RUN go build

CMD ["sh", "-c" ,"./greetclient"]