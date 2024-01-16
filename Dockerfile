FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on

RUN cd /build && git clone https://github.com/xerinox/go-htmx.git
RUN cd /build/go-htmx/main && go get && go build

EXPOSE 8080

WORKDIR /build/go-htmx/main

ENTRYPOINT [ "/build/go-htmx/main/main" ]
