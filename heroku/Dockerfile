FROM golang:latest as builder

COPY app/go.mod app/go.sum /go/src/app/

WORKDIR /go/src/app

RUN go mod download

COPY app /go/src/app

RUN apt-get update \
  && apt-get install -y git curl mariadb-client lsof\
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/* \
  && go get github.com/pressly/goose/v3/cmd/goose \
  && go build -o /go/bin/goose github.com/pressly/goose/v3/cmd/goose

RUN GOOS=linux GOARCH=amd64 go build -o /main .

FROM alpine:latest

COPY --from=builder /main .
RUN apk add --no-cache ca-certificates
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
# ADD https://github.com/golang/go/raw/master/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip

# RUN apk --update add tzdata && \
#     cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
#     apk del tzdata && \
#     rm -rf /var/cache/apk/*

CMD ./main
