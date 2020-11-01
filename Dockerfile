# Start from the latest golang base image
FROM golang:1.13 as builder

# Project maintainers
LABEL MAINTAINER="@mayoz <srcnckr@gmail.com>"

# Environments
ENV CGO_ENABLED 0
ENV GOOS linux

# Arguments
ARG PACKAGE_NAME

# Copy source files
COPY . /app

# Download dependencies via modules
WORKDIR /app
RUN go mod download

# Build the binary
WORKDIR /app/cmd/${PACKAGE_NAME}
RUN go build -o ./main .

# Running the build
FROM alpine:3.7

# Arguments
ARG PACKAGE_NAME

WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/cmd/${PACKAGE_NAME}/main /app/main

# Command to run the executable
CMD ["./main"]
