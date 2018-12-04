# cpu-model-nfd-plugin

[![Go Report Card](https://goreportcard.com/badge/github.com/ksimon1/cpu-model-nfd-plugin)](https://goreportcard.com/report/github.com/ksimon1/cpu-model-nfd-plugin)

**cpu-model-nfd-plugin** is plugin for [node-feature-discovery](https://github.com/kubernetes-sigs/node-feature-discovery). It creates list of all supported cpu models on host, which NFD then exposes as node labels.

**Usage:**
```
kubectl create -f cpu-model-nfd-plugin.yaml.template

kubectl describe nodes
```

This yaml file creates 1 pod. Pod contains NFD, libvirt (libvirt needs [Kubevirt](http://kubevirt.io/) to work properly), cpu-model-nfd-plugin.

**Description of [NFD pod](https://github.com/ksimon1/cpu-model-nfd-plugin/blob/master/cpu-model-nfd-plugin.yaml.template#L63):**

NFD pod contains 3 containers(NFD, libvirt, cpu-model-nfd-plugin). When cpu-model-nfd-plugin container is started, it copies cpu-model-nfd-plugin binary into NFD container in `/etc/kubernetes/node-feature-discovery/source.d/`. After it is copied, this container exits. In the libvirt container, output of virsh domcapabilities is saved into `/usr/share/virsh/` folder. This folder is shared between libvirt and NFD containers. After it is saved, this container exits. Now NFD container has all what it needs. It runs every 60 seconds cpu-model-nfd-plugin binary and this binary is taking data from `/usr/share/virsh/` folder. It parses data and prints data to stdout in format:
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
