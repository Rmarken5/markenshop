# Use the official Go image to build the app
FROM golang:1.20 AS builder

# Set the working directory to the app directory
WORKDIR /app

# Copy the source code into the container
COPY . .

WORKDIR /app/webhook

RUN ls

# Build the binary with debug information removed (for smaller size)
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o app  ./cmd/app.go

# Use a small base image to reduce the final image size
FROM scratch
# Copy the binary from the previous stage into this stage
COPY --from=builder /app/webhook/app /app

# Set the command to run when the container starts
ENTRYPOINT ["/app"]
