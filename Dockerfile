FROM registry.access.redhat.com/ubi8/ubi-minimal

MAINTAINER "Karel Šimon" <ksimon@redhat.com>

ENV container docker

RUN mkdir -p /plugin/dest
COPY kubevirt-cpu-nfd-plugin /plugin/dest
