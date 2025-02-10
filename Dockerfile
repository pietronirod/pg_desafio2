FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -trimpath -o /load-tester ./main.go

FROM scratch

COPY --from=builder /load-tester /load-tester

ENTRYPOINT ["/load-tester"]