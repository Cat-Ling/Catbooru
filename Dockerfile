# ---- Base ----
FROM golang:1.21-alpine AS base
WORKDIR /app
# Force the go compiler to use modules
ENV GO111MODULE=on
# We want to build a static binary
ENV CGO_ENABLED=0

# ---- Dependencies ----
FROM base AS deps
# Copy `go.mod` for definitions and `go.sum` for dependencies
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# ---- Build ----
FROM deps AS build
COPY . .
# Build the binary
RUN make build

# ---- Release ----
FROM gcr.io/distroless/static-debian11 AS release
WORKDIR /
COPY --from=build /app/booru-server /
COPY --from=build /app/config.yaml /
USER nonroot:nonroot
# Run the binary
ENTRYPOINT ["/booru-server"]
