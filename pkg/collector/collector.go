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
 * Copyright 2018 Red Hat, Inc.
 */

package collector

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"strings"
)

const usableNo string = "no"

// CollectData retrieves xml data from file and parse them.
// Output of this function is slice of usable lowercased cpu models.
// Only models with tag usable yes will be added to slice.
// <domainCapabilities>
//   <cpu>
//     <mode name='custom' supported='yes'>
//       <model usable='no'>EPYC-IBPB</model>
//       <model usable='yes'>Haswell</model>
//     </mode>
//   </cpu>
// </domainCapabilities>
// Output of this xml will be: ["haswell"]
func CollectData(hostDomCapabilitiesPath string, cpuModelBlackList map[string]bool) ([]string, error) {
	hostDomCapabilities := HostDomCapabilities{}
	err := getStructureFromFile(hostDomCapabilitiesPath, &hostDomCapabilities)
	if err != nil {
		return nil, err
	}
	cpus := make([]string, 0)

	for _, mode := range hostDomCapabilities.CPU.Mode {
		for _, model := range mode.Model {
			if model.Usable == usableNo || model.Usable == "" {
				continue
			}
			modelName := strings.ToLower(model.Name)
			if _, ok := cpuModelBlackList[modelName]; !ok {
				cpus = append(cpus, modelName)
			}
		}
	}
	return cpus, nil
}

//getStructureFromFile load data from file and unmarshals them into given structure
//Given structure has to be pointer
func getStructureFromFile(path string, structure interface{}) error {
	// Open xmlFile
	fileReader, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fileReader.Close()

	byteValue, err := ioutil.ReadAll(fileReader)
	if err != nil {
		return err
	}
	//unmarshal data into structure
	err = xml.Unmarshal(byteValue, structure)
	if err != nil {
		return err
	}
	return nil
}
