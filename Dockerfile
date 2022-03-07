FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN cd /build && git clone https://github.com/codecplyre/ascii-art-web.git

RUN cd /build/ascii-art-web && go build

EXPOSE 5500

ENTRYPOINT [ "/build/ascii-art-web/server" ]