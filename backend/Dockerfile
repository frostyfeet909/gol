FROM golang:1.25.6-bookworm AS build

ARG APP_NAME
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -trimpath -ldflags="-s -w" -o /out/api ./cmd/${APP_NAME}

FROM gcr.io/distroless/static-debian13:nonroot
COPY --from=build /out/api /api
ENTRYPOINT ["/api"]
