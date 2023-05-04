FROM golang:alpine AS build

RUN apk add --no-cache git build-base

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./

# todo add tests
# RUN CGO_ENABLED=1 go test -timeout 30s -v github.com/RomManager/server/tests

RUN mkdir /app
RUN CGO_ENABLED=1 go build \
    -installsuffix 'static' \
    -o /app/main \
    ./

FROM alpine as final

LABEL maintainer='pavelzw'
RUN addgroup -S nonroot \
    && adduser -S -u 10000 -g nonroot nonroot
COPY --from=build --chown=nonroot:nonroot /app /app
USER nonroot:nonroot
WORKDIR /app

ENTRYPOINT ["./main"]
