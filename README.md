# NFD-Host-supported-cpus

[![Go Report Card](https://goreportcard.com/badge/github.com/ksimon1/nfd-host-supported-cpus)](https://goreportcard.com/report/github.com/ksimon1/nfd-host-supported-cpus)

**NFD-Host-supported-cpus** is plugin for [node-feature-discovery](https://github.com/kubernetes-sigs/node-feature-discovery). It creates list of all supported cpu models on host, which NFD then exposes as node labels.

**Usage:**
```
kubectl create -f nfd-host-supported-cpus.yaml.template

kubectl describe nodes
```
**Result:**
![cpus](https://camo.githubusercontent.com/582985d780e4827856f862fbdd6b17f4f27f5c8c/68747470733a2f2f692e696d6775722e636f6d2f773643654343592e706e67)
