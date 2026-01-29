FROM golang:1.25.6-bookworm AS build

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@v1.16.6

COPY . .

# Turn off when I don't want swagger
RUN swag init -g ./cmd/api/swagger.go -o ./docs

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -tags=swagger -trimpath -ldflags="-s -w" -o /out/api ./cmd/api

FROM gcr.io/distroless/static-debian13:nonroot

COPY --from=build /out/api /api

ENTRYPOINT ["/api"]
