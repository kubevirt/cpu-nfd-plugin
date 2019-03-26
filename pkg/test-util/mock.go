/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2019 Red Hat, Inc.
 */

package testutil

var (
	DomainCapabilities = `<domainCapabilities>
  <cpu>
    <mode name='host-passthrough' supported='yes'/>
    <mode name='host-model' supported='yes'>
      <model fallback='allow'>Skylake-Client-IBRS</model>
      <vendor>Intel</vendor>
      <feature policy='require' name='ds'/>
      <feature policy='require' name='acpi'/>
      <feature policy='require' name='ss'/>
    </mode>
    <mode name='custom' supported='yes'>
      <model usable='no'>EPYC-IBPB</model>
      <model>fake-model-without-usable</model>
      <model usable='yes'>486</model>
      <model usable='yes'>Conroe</model>
      <model usable='yes'>coreduo</model>
      <model usable='yes'>IvyBridge</model>
      <model usable='yes'>Haswell</model>
    </mode>
  </cpu>
</domainCapabilities>`

	Features = []string{"apic", "clflush", "cmov"}

	CPUModelHaswellFeatures = `<cpus>
  <model name='Haswell'>
    <signature family='6' model='60'/>
    <vendor name='Intel'/>
    <feature name='aes'/>
    <feature name='apic'/>
    <feature name='bmi1'/>
    <feature name='clflush'/>
    <feature name='cmov'/>
  </model>
</cpus>`

	CPUModelIvyBridgeFeatures = `<cpus>
  <model name='IvyBridge'>
    <signature family='6' model='58'/>
    <vendor name='Intel'/>
    <feature name='aes'/>
    <feature name='apic'/>
    <feature name='clflush'/>
    <feature name='cmov'/>
  </model>
</cpus>
`
	CPUModelPenrynFeatures = `<cpus>
  <model name='Penryn'>
    <signature family='6' model='23'/>
    <vendor name='Intel'/>
    <feature name='apic'/>
    <feature name='clflush'/>
    <feature name='cmov'/>
  </model>
</cpus>`

	NewFeatures = []string{"bmi1", "aes"}

	CPUConfig = `obsoleteCPUs:
  - "486"
  - "pentium"
  - "pentium2"
  - "pentium3"
  - "pentiumpro"
  - "coreduo"
  - "n270"
  - "core2duo"
  - "Conroe"
  - "athlon"
  - "phenom"
minCPU:
  intel: "Penryn"
  amd: "Opteron_G1"`
)
