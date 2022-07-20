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

# FROM scratch

# COPY --from=builder /main ./

# ENV PORT=${PORT}
# ENTRYPOINT ["./main"]

FROM alpine:latest

COPY --from=builder /main /main

CMD /main $PORT
