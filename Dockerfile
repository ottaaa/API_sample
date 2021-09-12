FROM golang:1.17

WORKDIR /go/src/app

COPY . .

RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends git

RUN go get -u github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]
