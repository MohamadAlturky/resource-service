FROM golang:1.23.0-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
RUN ["ls"]

RUN go build -o . ./cmd/api/

EXPOSE 8082

CMD ["./api"]
