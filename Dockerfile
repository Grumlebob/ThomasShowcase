FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on

COPY go.mod /build
COPY go.sum /build/

RUN cd /build/ && git clone https://github.com/Grumlebob/ThomasShowcase.git
RUN cd /build/ThomasShowcase/server && go build ./...

EXPOSE 9080

ENTRYPOINT [ "/build/ThomasShowcase/server/server" ]