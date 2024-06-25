# Usar una imagen de Go oficial para construir la aplicación
FROM golang:1.21-alpine AS builder

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos del proyecto al contenedor
COPY . .

# Compilar la aplicación
RUN go mod tidy
RUN go build -o main ./cmd/server

# Crear una imagen mínima para ejecutar la aplicación
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

# Exponer el puerto en el que corre la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]
