# syntax=docker/dockerfile:1
FROM golang:1.22

# Install tzdata for timezones
RUN apk add --no-cache tzdata

# Set timezone (optional but useful)
ENV TZ=Asia/Bangkok

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the app
COPY . .

# Build the app
RUN go build -o stock_tracker .

# Expose the port
EXPOSE 3000

# Run the app
CMD ["./stock_tracker"]
