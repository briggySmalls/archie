## BUILD
FROM golang:alpine AS builder

# Install git for fetching the dependencies
RUN apk update && apk add --no-cache git

# Copy the repo here
WORKDIR /src
COPY . .

# Build and install the executable (in bin)
RUN go install ./cli/archie

# INSTALL
FROM scratch

# Copy in executable
COPY --from=builder /go/bin/archie .

# Expose port 80 for the API
EXPOSE 80

# Run the go program
ENTRYPOINT ["./archie"]

# Run on exposed port
CMD ["api --port 80"]
