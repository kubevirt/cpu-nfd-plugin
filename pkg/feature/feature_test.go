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
	"testing"

	testutil "github.com/ksimon1/kubevirt-cpu-nfd-plugin/pkg/test-util"
)

func TestGetPath(t *testing.T) {
	cpuName := "Penryn"
	path := GetPathCPUFefatures(cpuName)
	if path != "/etc/kubernetes/node-feature-discovery/source.d/cpu_map/x86_Penryn.xml" {
		t.Error("path should equal: /etc/kubernetes/node-feature-discovery/source.d/cpu_map/x86_Penryn.xml. Not to" + path)
		t.FailNow()
	}
}

func TestLoadFeatures(t *testing.T) {
	cpuName := "Penryn"
	LibvirtCPUMapFolder = "/tmp/"
	path := GetPathCPUFefatures(cpuName)

	err := testutil.WriteMockDataFile(path, testutil.CPUModelPenrynFeatures)
	if err != nil {
		t.Error("writeMockDataFile should not throw error: " + err.Error())
		t.FailNow()
	}
	features, err := LoadFeatures(cpuName)
	if err != nil {
		t.Error("LoadFeatures should not throw error: " + err.Error())
		t.FailNow()
	}

	for _, feature := range testutil.Features {
		if _, ok := features[feature]; !ok {
			t.Error("features should contain: " + feature + " feature")
			t.FailNow()
		}
	}

	testutil.DeleteMockFile(path)
}
