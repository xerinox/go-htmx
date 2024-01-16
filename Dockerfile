FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on

RUN go get github.com/xerinox/go-htmx/main
RUN cd /build && git clone https://github.com/xerinox/go-htmx.git

RUN cd /build/go-htmx/main && go build

EXPOSE 8080

ENTRYPOINT [ "/build/go-htmx/main/main" ]
