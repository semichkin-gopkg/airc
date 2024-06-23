FROM golang:1.22.2-alpine3.19

RUN apk update
RUN apk add git

COPY ./ /src
WORKDIR /src

RUN go build -o /usr/bin/airc ./cmd/airc/main.go
WORKDIR /

RUN rm -rf /src

ENTRYPOINT ["airc"]

CMD ["build-run", "-c", "/air/.air.toml"]