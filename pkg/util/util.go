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

package util

import (
	"kubevirt.io/kubevirt-cpu-nfd-plugin/pkg/file"
)

var ConfigPath = "/etc/kubernetes/node-feature-discovery/source.d/cpu-plugin-configmap.yaml"

func UnionMap(old, new map[string]bool) {
	for feature := range new {
		old[feature] = true
	}
}

func SubtractMap(a, b map[string]bool) map[string]bool {
	new := make(map[string]bool)
	for k := range a {
		if _, ok := b[k]; !ok {
			new[k] = true
		}
	}
	return new
}

func ConvertStringSliceToMap(s []string) map[string]bool {
	result := make(map[string]bool)
	for _, v := range s {
		result[v] = true
	}
	return result
}

func LoadConfig() (Config, error) {
	config := Config{}
	err := file.GetStructureFromYamlFile(ConfigPath, &config)
	if err != nil {
		return Config{}, err
	}
	return config, err
}

func (c *Config) GetObsoleteCPUMap() map[string]bool {
	return ConvertStringSliceToMap(c.ObsoleteCPUs)
}

func (c *Config) GetMinCPUByVendor(vendor string) string {
	var minCPU string
	if vendor == "Intel" {
		//set default baseline cpu model
		if c.MinCPU.Intel == "" {
			minCPU = "Conroe"
		} else {
			minCPU = c.MinCPU.Intel
		}
	} else {
		//set default baseline cpu model
		if c.MinCPU.AMD == "" {
			minCPU = "Opteron_G1"
		} else {
			minCPU = c.MinCPU.AMD
		}
	}

	return minCPU
}
