FROM registry.access.redhat.com/ubi8/ubi-minimal

MAINTAINER "The KubeVirt Project" <kubevirt-dev@googlegroups.com>

ENV container docker

RUN mkdir -p /plugin/dest
COPY cpu-nfd-plugin /plugin/dest
