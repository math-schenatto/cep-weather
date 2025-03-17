# Estágio de build
FROM golang:1.21 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o cloudrun

# Estágio de execução
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/cloudrun .

# Instala dependências necessárias (certificados CA e shell)
RUN apk --no-cache add ca-certificates tzdata

ENTRYPOINT ["./cloudrun"]