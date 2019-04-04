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

package feature

import (
	"github.com/ksimon1/kubevirt-cpu-nfd-plugin/pkg/file"
)

var LibvirtCPUMapFolder = "/etc/kubernetes/node-feature-discovery/source.d/cpu_map/"

//LoadFeatures loads features for given cpu name
func LoadFeatures(cpuModelName string) (map[string]bool, error) {
	if cpuModelName == "" {
		return map[string]bool{}, nil
	}

	cpuFeatures := FeatureModel{}
	cpuFeaturepath := GetPathCPUFefatures(cpuModelName)
	err := file.GetStructureFromXMLFile(cpuFeaturepath, &cpuFeatures)
	if err != nil {
		return nil, err
	}

	features := make(map[string]bool)
	for _, f := range cpuFeatures.Model.Features {
		features[f.Name] = true
	}
	return features, nil
}

//GetPathCPUFefatures creates path where folder with cpu models is
func GetPathCPUFefatures(name string) string {
	return LibvirtCPUMapFolder + "x86_" + name + ".xml"
}
