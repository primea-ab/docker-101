# Buildtime
FROM golang:1.21 AS builder

WORKDIR /work

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o binary ./cmd

# Runtime
FROM alpine:latest

COPY --from=builder /work/binary /binary

CMD ["/binary"]