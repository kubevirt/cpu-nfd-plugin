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
	"strconv"
	"testing"

	"kubevirt.io/kubevirt-cpu-nfd-plugin/pkg/feature"
	testutil "kubevirt.io/kubevirt-cpu-nfd-plugin/pkg/test-util"
)

func prepareFiles(t *testing.T) {
	domCapabilitiesFilePath := "/tmp/virsh-domcapabilities.xml"
	err := testutil.WriteMockDataFile(domCapabilitiesFilePath, testutil.DomainCapabilities)
	if err != nil {
		t.Error("writeMockDataFile should not throw error: " + err.Error())
		t.FailNow()
	}
	feature.LibvirtCPUMapFolder = "/tmp/"
	penrynPath := feature.GetPathCPUFefatures("Penryn")
	err = testutil.WriteMockDataFile(penrynPath, testutil.CPUModelPenrynFeatures)
	if err != nil {
		t.Error("writeMockDataFile should not throw error: " + err.Error())
		t.FailNow()
	}
	ivyBridgePath := feature.GetPathCPUFefatures("IvyBridge")
	err = testutil.WriteMockDataFile(ivyBridgePath, testutil.CPUModelIvyBridgeFeatures)
	if err != nil {
		t.Error("writeMockDataFile should not throw error: " + err.Error())
		t.FailNow()
	}
	haswellPath := feature.GetPathCPUFefatures("Haswell")
	err = testutil.WriteMockDataFile(haswellPath, testutil.CPUModelHaswellFeatures)
	if err != nil {
		t.Error("writeMockDataFile should not throw error: " + err.Error())
		t.FailNow()
	}
}

func deleteFiles() {
	feature.LibvirtCPUMapFolder = "/tmp/"
	testutil.DeleteMockFile(feature.GetPathCPUFefatures("Penryn"))
	testutil.DeleteMockFile(feature.GetPathCPUFefatures("Haswell"))
	testutil.DeleteMockFile(feature.GetPathCPUFefatures("IvyBridge"))
	testutil.DeleteMockFile(feature.GetPathCPUFefatures(domCapabilitiesFilePath))
}

// TestCollectData tests CollectData function
func TestCollectData(t *testing.T) {
	domCapabilitiesFilePath = "/tmp/virsh-domcapabilities.xml"
	feature.LibvirtCPUMapFolder = "/tmp/"
	prepareFiles(t)
	cpuModels, cpuFeatures, err := CollectData()
	if err != nil {
		t.Error("CollectData should not throw error: " + err.Error())
	}

	if len(cpuModels) != 2 {
		t.Error("CollectData should return 2 cpu models, it returns: " + strconv.Itoa(len(cpuModels)))
	}

	if len(cpuFeatures) != 2 {
		t.Error("CollectData should return 2 cpu features, it returns: " + strconv.Itoa(len(cpuFeatures)))
	}

	for _, feature := range testutil.NewFeatures {
		if _, ok := cpuFeatures[feature]; !ok {
			t.Error("features should contain: " + feature + " feature")
			t.FailNow()
		}
	}

	deleteFiles()
}
