# NFD-Host-supported-cpus

[![Go Report Card](https://goreportcard.com/badge/github.com/ksimon1/nfd-host-supported-cpus)](https://goreportcard.com/report/github.com/ksimon1/nfd-host-supported-cpus)

**NFD-Host-supported-cpus** is plugin for [node-feature-discovery](https://github.com/kubernetes-sigs/node-feature-discovery). It creates list of all supported cpu models on host, which NFD then exposes as node labels.

**Usage:**
```
kubectl create -f nfd-host-supported-cpus.yaml.template

kubectl describe nodes
```
