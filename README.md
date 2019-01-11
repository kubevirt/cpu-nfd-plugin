# cpu-model-nfd-plugin

[![Go Report Card](https://goreportcard.com/badge/github.com/ksimon1/cpu-model-nfd-plugin)](https://goreportcard.com/report/github.com/ksimon1/cpu-model-nfd-plugin)

**cpu-model-nfd-plugin** is plugin for [cluster-nfd-operator](https://github.com/openshift/cluster-nfd-operator). It creates list of all supported cpu models on host, which NFD then exposes as node labels.

**Usage:**
```
kubectl create -f cpu-model-nfd-plugin.yaml

kubectl describe nodes
```

This yaml file creates 1 pod. Pod contains libvirt (libvirt needs [Kubevirt](http://kubevirt.io/) to work properly), cpu-model-nfd-plugin.

**Description of [NFD pod](https://github.com/ksimon1/cpu-model-nfd-plugin/blob/master/cpu-model-nfd-plugin.yaml):**

NFD pod contains 2 containers(libvirt, cpu-model-nfd-plugin). In the libvirt container, output of virsh domcapabilities is saved into `/host-hook/virsh_domcapabilities.xml` folder. This folder is shared between libvirt and NFD container. After it is saved, this container exits. NFD runs every 60 seconds cpu-model-nfd-plugin binary. It parses data and prints data to stdout in format:
```
/cpu-model-haswell
/cpu-model-core2duo
...
```
NFD takes this output and creates labels from them.

**CPU model black list:**

CPU model black list can be set for this plugin. Just write all models, which are not needed into env variables in node-feature-discovery container in format `value: "model1 model2 modeln"`

```
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: node-feature-discovery
spec:
    spec:
      containers:
        name: node-feature-discovery
        - env:
          - name: CPU_MODEL_BLACK_LIST
            value: "opteron_g1 opteron_g2 haswell"
```
**Result:**
![cpus](https://camo.githubusercontent.com/582985d780e4827856f862fbdd6b17f4f27f5c8c/68747470733a2f2f692e696d6775722e636f6d2f773643654343592e706e67)
