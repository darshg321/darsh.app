FROM golang:1.27-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /term .

FROM alpine:3
COPY --from=build /term /term
# mount a volume here so the host key survives redeploys (otherwise visitors get a "host key changed" warning)
ENV HOST_KEY=/data/id_ed25519 PORT=22
EXPOSE 22
ENTRYPOINT ["/term"]
