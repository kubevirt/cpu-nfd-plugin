# kubevirt-cpu-nfd-plugin

**kubevirt-cpu-nfd-plugin** is plugin for [kubevirt-node-labeller](https://github.com/ksimon1/kubevirt-cpu-node-labeller). It creates list of all supported cpu models and features on host, which cpu-node-labeller then exposes as node labels.

**Usage:**
```
oc apply -f cpu-model-labeller-plugin.yaml
```

**How it works**

Plugin parses libvirt data (cpu models, features) and prints data to stdout in format:
```
/cpu-model-haswell
/cpu-model-core2duo
/cpu-feature-aes
...
```
cpu-node-labeller takes this output and creates labels from them. Plugin has list of unsupported old cpus, which are not in the label list: 486, pentium, pentium2, pentium3, pentiumpro, coreduo, n270, core2duo, Conroe, athlon, phenom. As basic models for Intel is cpu model Penryn and for AMD it is Opteron_G1. Features which have these cpus are taken as basic features. These basic features are not in the label list. Feature labels are created as subtraction between set of newer cpu features and set of basic cpu features, e.g.:
Haswell has: aes, apic, clflush
Penryr has: apic, clflush
subtraction is: aes. So label will be created only with aes feature.

**Result:**
![cpus](https://camo.githubusercontent.com/582985d780e4827856f862fbdd6b17f4f27f5c8c/68747470733a2f2f692e696d6775722e636f6d2f773643654343592e706e67)
