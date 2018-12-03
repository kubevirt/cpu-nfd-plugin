FROM fedora:28

MAINTAINER "Karel Šimon" <ksimon@redhat.com>

ENV container docker

RUN mkdir -p /plugin/dest
COPY cpu-model-nfd-plugin /plugin/dest
