# FROM golang:1.24-alpine

# WORKDIR /app

# RUN go install github.com/air-verse/air@latest

# COPY go.mod go.sum ./
# RUN go mod download

# COPY . .

# RUN go build -o main main.go

# CMD ["air", "-c", ".air.toml"]


FROM golang:1.24-alpine

WORKDIR /app

# Install necessary packages and air
RUN apk add --no-cache git
RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
