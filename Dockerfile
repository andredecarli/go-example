FROM golang:1.24.4-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o app ./cmd


FROM alpine:latest

WORKDIR /app

COPY --from=build /app/app /app/.env .

EXPOSE 8080

CMD ["./app"]
