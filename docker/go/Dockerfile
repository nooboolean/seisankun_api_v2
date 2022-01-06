FROM golang:latest

COPY app/go.mod app/go.sum /go/src/app/

WORKDIR /go/src/app

RUN go mod download

COPY app /go/src/app

RUN apt-get update \
  && apt-get install -y git mariadb-client lsof\
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/* \
  && go install github.com/cosmtrek/air@latest \
  && go get github.com/pressly/goose/v3/cmd/goose \
  && go build -o /go/bin/goose github.com/pressly/goose/v3/cmd/goose

EXPOSE 3000
CMD ["air", "-c", ".air.toml"]
