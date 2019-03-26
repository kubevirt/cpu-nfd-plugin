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

package collector

import (
	"kubevirt.io/kubevirt-cpu-nfd-plugin/pkg/feature"
	"kubevirt.io/kubevirt-cpu-nfd-plugin/pkg/file"
	"kubevirt.io/kubevirt-cpu-nfd-plugin/pkg/util"
)

const (
	usableNo string = "no"
)

var domCapabilitiesFilePath = "/etc/kubernetes/node-feature-discovery/source.d/virsh_domcapabilities.xml"

// CollectData retrieves xml data from file and parse them.
// Output of this function is slice of usable cpu models and features.
// Only models with tag usable yes will be used.
func CollectData() ([]string, map[string]bool, error) {
	hostDomCapabilities := HostDomCapabilities{}
	err := file.GetStructureFromXMLFile(domCapabilitiesFilePath, &hostDomCapabilities)
	if err != nil {
		return nil, nil, err
	}
	var config util.Config
	config, err = util.LoadConfig()
	if err != nil {
		return nil, nil, err
	}

	obsoleteCPUsx86 := config.GetObsoleteCPUMap()

	basicFeaturesMap := make(map[string]bool)
	cpus := make([]string, 0)
	features := make(map[string]bool)

	for _, mode := range hostDomCapabilities.CPU.Mode {
		if mode.Vendor.Name != "" {
			minCPU := config.GetMinCPUByVendor(mode.Vendor.Name)
			var err error
			basicFeaturesMap, err = getBasicFeatures(minCPU)
			if err != nil {
				return nil, nil, err
			}

		}
		for _, model := range mode.Model {
			if _, ok := obsoleteCPUsx86[model.Name]; ok || model.Usable == usableNo || model.Usable == "" {
				continue
			}

			newFeatures, _ := parseFeatures(basicFeaturesMap, model.Name)
			util.UnionMap(features, newFeatures)

			cpus = append(cpus, model.Name)
		}
	}
	return cpus, features, nil
}

func getBasicFeatures(minCPU string) (map[string]bool, error) {
	fm, err := feature.LoadFeatures(minCPU)
	if err != nil {
		return nil, err
	}
	return fm, nil
}

func parseFeatures(basicFeatures map[string]bool, cpuName string) (map[string]bool, error) {
	features, err := feature.LoadFeatures(cpuName)
	if err != nil {
		return nil, err
	}
	return util.SubtractMap(features, basicFeatures), nil
}
