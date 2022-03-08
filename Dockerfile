FROM golang:alpine

RUN mkdir /ascii-art-web

ADD . /ascii-art-web

WORKDIR /ascii-art-web

RUN go build -o main .

EXPOSE 5500

CMD ["/ascii-art-web/main"]