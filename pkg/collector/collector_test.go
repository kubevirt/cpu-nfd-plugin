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
	"fmt"
	"strconv"
	"testing"

	"kubevirt.io/cpu-nfd-plugin/pkg/config"
	"kubevirt.io/cpu-nfd-plugin/pkg/feature"
	testutil "kubevirt.io/cpu-nfd-plugin/pkg/test-util"
)

func TestCollectData(t *testing.T) {

	var testCases = []struct {
		testCase         string
		prepareDataFiles func(*testing.T)
		checks           func(*testing.T, []string, map[string]bool, error)
	}{
		{
			testCase: "Everything is ok",
			prepareDataFiles: func(t *testing.T) {
				prepareFileDomCapabilities(t)
				prepareFilesFeatures(t)
				prepareConfigFile(t)
			},
			checks: func(t *testing.T, cpuModels []string, cpuFeatures map[string]bool, err error) {
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
			},
		},
		{
			testCase: "Config is not ready",
			prepareDataFiles: func(t *testing.T) {
				prepareFileDomCapabilities(t)
				prepareFilesFeatures(t)
			},
			checks: func(t *testing.T, cpuModels []string, cpuFeatures map[string]bool, err error) {
				if err != nil {
					t.Error("CollectData should not throw error: " + err.Error())
				}

				if len(cpuModels) != 5 {
					t.Error("CollectData should return 5 cpu models, it returns: " + strconv.Itoa(len(cpuModels)))
				}

				if len(cpuFeatures) != 5 {
					t.Error("CollectData should return 5 cpu features, it returns: " + strconv.Itoa(len(cpuFeatures)))
				}
			},
		},
		{
			testCase: "Dom capabilities file is not ready",
			prepareDataFiles: func(t *testing.T) {
				prepareFilesFeatures(t)
			},
			checks: func(t *testing.T, cpuModels []string, cpuFeatures map[string]bool, err error) {
				if err == nil {
					t.Error("CollectData should throw error: ")
				}
			},
		},
		{
			testCase: "No cpu model is usable",
			prepareDataFiles: func(t *testing.T) {
				prepareFileDomCapabilitiesNothingUsable(t)
				prepareFilesFeatures(t)
			},
			checks: func(t *testing.T, cpuModels []string, cpuFeatures map[string]bool, err error) {
				if err != nil {
					t.Error("CollectData should not throw error: " + err.Error())
				}

				if len(cpuModels) != 0 {
					t.Error("CollectData should return 0 cpu models, it returns: " + strconv.Itoa(len(cpuModels)))
				}

				if len(cpuFeatures) != 0 {
					t.Error("CollectData should return 0 cpu features, it returns: " + strconv.Itoa(len(cpuFeatures)))
				}
			},
		},
	}

	for _, testCase := range testCases {
		fmt.Println("Running test case: " + testCase.testCase)
		testCase.prepareDataFiles(t)
		//set all path to /tmp
		domCapabilitiesFilePath = "/tmp/virsh-domcapabilities.xml"
		feature.LibvirtCPUMapFolder = "/tmp/"
		config.ConfigPath = "/tmp/cpu-plugin-configmap.yaml"

		cpuModels, cpuFeatures, err := CollectData()
		testCase.checks(t, cpuModels, cpuFeatures, err)
		deleteFiles()
	}
}

func prepareFileDomCapabilities(t *testing.T) {
	domCapabilitiesFilePath := "/tmp/virsh-domcapabilities.xml"
	err := testutil.WriteMockDataFile(domCapabilitiesFilePath, testutil.DomainCapabilities)
	if err != nil {
		t.Error("writeMockDataFile should not throw error: " + err.Error())
		t.FailNow()
	}
}

func prepareFileDomCapabilitiesNothingUsable(t *testing.T) {
	domCapabilitiesFilePath := "/tmp/virsh-domcapabilities.xml"
	err := testutil.WriteMockDataFile(domCapabilitiesFilePath, testutil.DomainCapabilitiesNothingUsable)
	if err != nil {
		t.Error("writeMockDataFile should not throw error: " + err.Error())
		t.FailNow()
	}
}

func prepareFilesFeatures(t *testing.T) {
	feature.LibvirtCPUMapFolder = "/tmp/"
	penrynPath := feature.GetPathCPUFefatures("Penryn")
	err := testutil.WriteMockDataFile(penrynPath, testutil.CPUModelPenrynFeatures)
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

func prepareConfigFile(t *testing.T) {
	cpuConfigPath := "/tmp/cpu-plugin-configmap.yaml"
	err := testutil.WriteMockDataFile(cpuConfigPath, testutil.CPUConfig)
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
	testutil.DeleteMockFile("/tmp/virsh-domcapabilities.xml")
	testutil.DeleteMockFile("/tmp/cpu-plugin-configmap.yaml")
}
