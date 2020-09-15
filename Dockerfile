FROM golang:1.14-buster AS builder
WORKDIR /go/src/devfarm
COPY . .
RUN go build -a -tags netgo -installsuffix netgo -v -o ./dist/devfarm ./cmd/devfarm/main.go


FROM amazon/aws-cli:2.0.48
COPY --from=builder /go/src/devfarm/dist/* /usr/local/bin/
VOLUME ["/app"]

ENTRYPOINT ["devfarm"]
