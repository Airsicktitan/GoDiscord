# Use Go official image to build AND run
FROM golang:1.22.3

WORKDIR /app

# Copy and build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o bot

# Set env and run
ENV TOKEN={TokenHere}
CMD ["./bot"]
