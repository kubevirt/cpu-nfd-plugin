FROM golang:1.10 AS builder

MAINTAINER "The KubeVirt Project" <kubevirt-dev@googlegroups.com>

WORKDIR /go/src/kubevirt.io/cpu-nfd-plugin

ENV GOPATH=/go

COPY . .

RUN go test ./...

RUN CGO_ENABLED=0 GOOS=linux go build -o /cpu-nfd-plugin cmd/cpu-nfd-plugin/cpu-nfd-plugin.go

FROM registry.access.redhat.com/ubi8/ubi-minimal

RUN mkdir -p /plugin/dest

COPY --from=builder /cpu-nfd-plugin /plugin/dest/cpu-nfd-plugin
