FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main

FROM scratch

COPY --from=builder /app/main /
COPY app.ini /

EXPOSE 8080

CMD ["/main"] 