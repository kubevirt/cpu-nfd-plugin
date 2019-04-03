# kubevirt-cpu-nfd-plugin

**kubevirt-cpu-nfd-plugin** is plugin for [kubevirt-node-labeller](https://github.com/ksimon1/kubevirt-cpu-node-labeller). It creates list of all supported cpu models and features on host, which cpu-node-labeller then exposes as node labels.

**How it works**

Plugin parses libvirt data (cpu models, features) and prints data to stdout in format:
```
/cpu-model-haswell
/cpu-model-core2duo
/cpu-feature-aes
...
```
cpu-node-labeller takes this output and creates labels from them. 

The plugin can hide old cpus. This can be done by creating config map:
```
apiVersion: v1
kind: ConfigMap
metadata:
  name: kubevirt-cpu-plugin-configmap
data:
  cpu-plugin-configmap.yaml: |- 
    obsoleteCPUs:
      - "486"
      - "pentium"
      - "pentium2"
      - "pentium3"
      - "pentiumpro"
    minCPU: "Penryn"
```
This config map has to be created before kubevirt-node-labeller is created (example of deploy manifest can be found [here](https://github.com/ksimon1/kubevirt-node-labeller/blob/master/kubevirt-node-labeller.yaml)), otherwise plugin will show all cpu models. Plugin will not reload when config map is changed.

Add cpu model into obsoleteCPUs array and cpu model will not be listed in labels.
User can define minimal cpu model. This model's features will be used as basic features. These basic features are not in the label list. Feature labels are created as subtraction between set of newer cpu features and set of basic cpu features, e.g.:
Haswell has: aes, apic, clflush
Penryr has: apic, clflush
subtraction is: aes. So label will be created only with aes feature.

