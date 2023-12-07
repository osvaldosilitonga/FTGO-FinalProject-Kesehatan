FROM golang:1.21-alpine AS build

WORKDIR /app
COPY . .

RUN go mod download && \
    go mod verify && \
    CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Runtime
FROM alpine:3.18

WORKDIR /app
COPY .env .
COPY template ./template
COPY --from=build /app/main .

# EXPOSE 8081

CMD ["./main"]