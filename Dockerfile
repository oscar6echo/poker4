

FROM golang:1.19-alpine AS builder

WORKDIR /go/src/

COPY ./poker-go .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o app .


####################


FROM scratch

WORKDIR /

COPY --from=builder /go/src/app /

ENV GIN_MODE=release

EXPOSE 5000

CMD ["/app"]
