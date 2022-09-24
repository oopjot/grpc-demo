FROM golang:alpine AS builder
RUN apk update && apk add git

WORKDIR /build

COPY . .

RUN go mod init github.com/oopjot/grpc-demo/adder
RUN go mod tidy
RUN go mod verify

RUN CGO_ENABLED=0 go build -o ./adder-bin

FROM scratch
WORKDIR /
COPY --from=builder /build/adder-bin .
EXPOSE 50000
ENTRYPOINT ["/adder-bin"]

