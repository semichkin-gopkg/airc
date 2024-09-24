FROM golang:1.23.1-alpine3.20

RUN apk update
RUN apk add git
RUN apk add curl

COPY ./ /src
WORKDIR /src

RUN go build -o /usr/bin/airc ./cmd/airc/main.go
WORKDIR /

RUN rm -rf /src

ENTRYPOINT ["airc"]

CMD ["build-run", "-c", "/air/.air.toml"]