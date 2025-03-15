FROM golang:1.24-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o analytics-api .
FROM gcr.io/distroless/base
WORKDIR /app
COPY --from=builder /app/analytics-api .
CMD ["./analytics-api"]