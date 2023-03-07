ARG  DISTROLESS_IMAGE=gcr.io/distroless/static:nonroot
ARG AppListen=80

# Use the official Go image as the base image
FROM golang:alpine3.17 AS build

# Set the working directory to /app
WORKDIR /app

# Install deps
RUN apk update --no-cache && apk add pkgconf openssl-dev gcc libc-dev

# Copy the source code to the container
COPY . .

ENV GO111MODULE=on
RUN go mod download

# Build the binary
RUN GOOS=linux GOARCH=amd64 go build -tags static -o /go/bin/chatsapi .

# Use a lightweight image for the final stage
FROM ${DISTROLESS_IMAGE}

USER 65532:65532

# Copy the binary from the previous stage
COPY --from=build /go/bin/chatsapi /go/bin/chatsapi


# Expose the port that the application will listen on
EXPOSE ${AppListen}

# Run the binary
CMD ["/go/bin/chatsapi"]
