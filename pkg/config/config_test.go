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

package config

import (
	"reflect"
	"testing"
)

func TestGetObsoleteCPUMap(t *testing.T) {
	var testCases = []struct {
		conf   Config
		result map[string]bool
	}{
		{
			conf: Config{
				ObsoleteCPUs: []string{"Conroe", "Haswell", "Penryn"},
			},
			result: map[string]bool{
				"Conroe":  true,
				"Haswell": true,
				"Penryn":  true,
			},
		},
		{
			conf: Config{
				ObsoleteCPUs: []string{},
			},
			result: map[string]bool{},
		},
	}
	for _, testCase := range testCases {
		m := testCase.conf.GetObsoleteCPUMap()
		if !reflect.DeepEqual(m, testCase.result) {
			t.Error(m, " should equal: ", testCase.result)
		}
	}
}

func TestGetMinCPUByVendor(t *testing.T) {
	var testCases = []struct {
		conf     Config
		provider string
		result   string
	}{
		{
			conf: Config{
				MinCPU: "Penryn",
			},
			result: "Penryn",
		},
		{
			conf: Config{
				MinCPU: "",
			},
			result: "",
		},
	}
	for _, testCase := range testCases {
		result := testCase.conf.GetMinCPU()
		if result != testCase.result {
			t.Error(result, " should equal: ", testCase.result)
		}
	}
}
