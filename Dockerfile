FROM golang:1.4.2

RUN mkdir -p /go/src/app
WORKDIR /go/src

RUN go get github.com/revel/revel
RUN go get github.com/revel/cmd/revel

COPY . /go/src/app

EXPOSE 9000
CMD ["revel", "run", "app", "prod"]