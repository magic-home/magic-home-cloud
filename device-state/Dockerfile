FROM golang:1.9.4 as build
  
WORKDIR /go/src/github.com/magic-home/magic-home-cloud/device-state

COPY vendor         vendor
COPY server.go      .

# Run a gofmt and exclude all vendored code.
RUN test -z "$(gofmt -l $(find . -type f -name '*.go' -not -path "./vendor/*"))" \
    && go test $(go list ./... | grep -v integration | grep -v /vendor/ | grep -v /template/) -cover \
    && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o device-state .

FROM alpine:3.7

RUN addgroup -S service \
    && adduser -S -g service service

WORKDIR /home/service

EXPOSE 3312

COPY --from=build /go/src/github.com/magic-home/magic-home-cloud/device-state/device-state .

RUN chown -R service:service ./

USER service

CMD ["./device-state"]
