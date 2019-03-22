# kubevirt-cpu-nfd-plugin

**kubevirt-cpu-nfd-plugin** is plugin for [kubevirt-node-labeller](https://github.com/ksimon1/kubevirt-cpu-node-labeller). It creates list of all supported cpu models on host, which cpu-node-labeller then exposes as node labels.

**Usage:**
```
oc apply -f cpu-model-labeller-plugin.yaml
```

**How it works**

In the libvirt container, output of virsh domcapabilities is saved into `/host-hook/virsh_domcapabilities.xml` folder. This folder is shared between libvirt and cpu-node-labeller container. Plugin parses libvirt data and prints data to stdout in format:
```
/cpu-model-haswell
/cpu-model-core2duo
...
```
cpu-node-labeller takes this output and creates labels from them.

**Result:**
![cpus](https://camo.githubusercontent.com/582985d780e4827856f862fbdd6b17f4f27f5c8c/68747470733a2f2f692e696d6775722e636f6d2f773643654343592e706e67)
