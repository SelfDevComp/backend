# -- BUILD STAGE --
FROM golang:1.26-alpine AS build-stage

WORKDIR /app

COPY ../backend/go.mod ../backend/go.sum ./
RUN go mod download
COPY ../backend/. .
RUN CGO_ENABLED=0 GOOS=linux go build -o self-dev ./cmd/api

# -- RUN STAGE --
FROM golang:1.26-alpine 

WORKDIR /app

COPY ../infra/.env /app
COPY --from=build-stage /app /app

EXPOSE 8080

ENTRYPOINT ["/app/self-dev"]