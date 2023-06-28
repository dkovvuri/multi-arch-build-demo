# Use the official Go image as the base image
FROM golang:1.17-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download the dependencies
COPY go.mod ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o mywebapp main.go


# Building the application container
FROM scratch

# Copying the binary from the compiled binary
COPY --from=build /app/mywebapp .

# Expose the port that the application listens on
EXPOSE 8080

CMD ["./mywebapp"]