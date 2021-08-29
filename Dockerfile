FROM golang AS builder
WORKDIR /app

COPY ./ ./

RUN go build cmd/main.go

FROM golang:1.14
WORKDIR /go/bin

COPY --from=builder /app/main /go/bin
COPY --from=builder /app/configs /go/bin/configs
COPY --from=builder /app/src /go/bin/src

CMD [ "/go/bin/main" ]
