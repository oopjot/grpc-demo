FROM golang:alpine AS builder
RUN apk update && apk add git

WORKDIR /build

COPY . .

RUN go mod init github.com/oopjot/grpc-demo/fibonacci
RUN go mod tidy
RUN go mod verify

RUN CGO_ENABLED=0 go build -o ./fibonacci-bin

FROM scratch
WORKDIR /
COPY --from=builder /build/fibonacci-bin .
EXPOSE 50001
ENTRYPOINT ["/fibonacci-bin"]

