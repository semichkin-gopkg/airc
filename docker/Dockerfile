FROM golang:1.22.2-alpine3.19

RUN apk update
RUN apk add git

RUN go install github.com/semichkin-gopkg/airc/cmd/airc@v0.0.7

ENTRYPOINT ["airc"]

CMD ["build-run", "-c", "/air/.air.toml"]