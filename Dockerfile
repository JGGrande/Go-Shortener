#Etapa de build do projeto
FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git

WORKDIR /src

COPY . .

RUN go mod download
RUN go build -buildvcs=false -o main.exe .

#Etapa de execução
FROM alpine:latest

COPY --from=builder src/main.exe .

ENTRYPOINT [ "./main.exe" ]