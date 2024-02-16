############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS go_service
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR /src
COPY . .
# Fetch dependencies.
# Using go get.
RUN go mod download
# Build the binary.
RUN go build -ldflags="-w -s" -o TransactionsFromFileService ./cmd/.

ENTRYPOINT ["/src/TransactionsFromFileService"]


