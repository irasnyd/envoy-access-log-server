# syntax=docker/dockerfile:1

FROM golang:1.18-buster AS build

WORKDIR /app
COPY . .
RUN ls -lR ./
RUN go build -o /envoy-access-log-server

FROM gcr.io/distroless/base-debian10
WORKDIR /
COPY --from=build /envoy-access-log-server /envoy-access-log-server
EXPOSE 9001
USER nonroot:nonroot
ENTRYPOINT ["/envoy-access-log-server"]
