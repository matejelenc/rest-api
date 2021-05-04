#Build stage
 FROM golang:alpine AS builder
 WORKDIR /app
 COPY . .
 RUN go build -o main main.go

 #Run stage
 FROM alpine
 WORKDIR /app
 COPY --from=builder /app/main .
 COPY .env .

 EXPOSE 8080
 CMD [ "/app/main" ]