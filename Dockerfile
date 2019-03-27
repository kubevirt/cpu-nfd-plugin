FROM centos:7

MAINTAINER "Karel Šimon" <ksimon@redhat.com>

ENV container docker

RUN yum -y update && yum clean all

RUN mkdir -p /plugin/dest
COPY kubevirt-cpu-nfd-plugin /plugin/dest
