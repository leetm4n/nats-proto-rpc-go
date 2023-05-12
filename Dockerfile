FROM golang as builder

WORKDIR /go/src/nats-proto-rpc-go

ENV GOOS=linux GOARCH=amd64 CGO_ENABLED=0

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o protoc-gen-natsrpcgo ./cmd/protoc-gen-natsrpcgo

FROM scratch

COPY --from=builder /go/src/nats-proto-rpc-go/protoc-gen-natsrpcgo /protoc-gen-natsrpcgo

ENTRYPOINT ["/protoc-gen-natsrpcgo"]
