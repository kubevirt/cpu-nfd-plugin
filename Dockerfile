FROM fedora:28

MAINTAINER "Karel Šimon" <ksimon@redhat.com>

ENV container docker

COPY nfd-host-cpus /usr/sbin/nfd-host-cpus
