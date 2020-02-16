FROM golang:1.13.3

WORKDIR /go/src/app

COPY . .

ENV GO111MODULE on
ENV SERVER "greetserver"

RUN go build

CMD ["sh", "-c" ,"./greetclient"]