FROM golang:alpine AS builder
RUN apk update && apk add git

WORKDIR /build

COPY . .

RUN go mod init github.com/oopjot/grpc-demo/gateway
RUN go mod tidy
RUN go mod verify

RUN CGO_ENABLED=0 go build -o ./gateway-bin

FROM scratch
WORKDIR /
COPY --from=builder /build/gateway-bin .
EXPOSE 8000
ENTRYPOINT ["/gateway-bin"]

