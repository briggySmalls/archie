## BUILD
FROM golang:alpine AS builder

# Install git for fetching the dependencies
RUN apk update && apk add --no-cache git make

# Set working diretory
WORKDIR /src

# Copy the go mod files and download dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the rest of the source files
COPY . .

# Build and install the executable (in bin)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install ./cli/archie

# INSTALL
FROM scratch

# Copy in executable
COPY --from=builder /go/bin/archie /app/archie

# Expose port 80 for the API
EXPOSE 80

# Run the go program
ENTRYPOINT ["/app/archie"]

# Run on exposed port
CMD ["api", "--port", "80"]
