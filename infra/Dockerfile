# -- BUILD STAGE --
FROM golang:1.26-alpine AS build-stage

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o self-dev ./cmd/api

# -- RUN STAGE --
FROM golang:1.26-alpine 

WORKDIR /app

COPY --from=build-stage /app /app
COPY --from=build-stage /go/bin/goose /usr/local/bin/goose

EXPOSE 8080

ENTRYPOINT ["/app/self-dev"]