FROM golang:latest

COPY app /go/src/app

WORKDIR /go/src/app

RUN go mod download

RUN apt-get update \
  && apt-get install -y git mariadb-client lsof\
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/* \
  && go build -o /go/bin/air github.com/cosmtrek/air

EXPOSE 3000
CMD ["air", "-c", ".air.toml"]
