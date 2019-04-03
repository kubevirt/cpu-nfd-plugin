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
	"reflect"
	"testing"
)

func TestUnionMap(t *testing.T) {
	var testCases = []struct {
		inA    map[string]bool
		inB    map[string]bool
		result map[string]bool
	}{
		{
			inA: map[string]bool{
				"a": true,
				"b": true,
			},
			inB: map[string]bool{
				"c": true,
			},
			result: map[string]bool{
				"a": true,
				"b": true,
				"c": true,
			},
		},
		{
			inA: map[string]bool{},
			inB: map[string]bool{
				"c": true,
			},
			result: map[string]bool{
				"c": true,
			},
		},
		{
			inA: map[string]bool{
				"a": true,
			},
			inB: map[string]bool{},
			result: map[string]bool{
				"a": true,
			},
		},
	}
	for _, testCase := range testCases {
		result := UnionMap(testCase.inA, testCase.inB)
		if !reflect.DeepEqual(result, testCase.result) {
			t.Error(result, " should equal: ", testCase.result)
		}
	}
}
func TestSubtractMap(t *testing.T) {

	var testCases = []struct {
		inA    map[string]bool
		inB    map[string]bool
		result map[string]bool
	}{
		{
			inA: map[string]bool{
				"a": true,
				"b": true,
			},
			inB: map[string]bool{
				"b": true,
			},
			result: map[string]bool{
				"a": true,
			},
		},
		{
			inA: map[string]bool{
				"a": true,
				"b": true,
			},
			inB: map[string]bool{},
			result: map[string]bool{
				"a": true,
				"b": true,
			},
		},
		{
			inA: map[string]bool{},
			inB: map[string]bool{
				"a": true,
				"b": true,
			},
			result: map[string]bool{},
		},
	}
	for _, testCase := range testCases {
		result := SubtractMap(testCase.inA, testCase.inB)
		if !reflect.DeepEqual(result, testCase.result) {
			t.Error(result, " should equal: ", testCase.result)
		}
	}
}

func TestConvertStringSliceToMap(t *testing.T) {
	var testCases = []struct {
		inA    []string
		result map[string]bool
	}{
		{
			inA: []string{"a", "b"},
			result: map[string]bool{
				"a": true,
				"b": true,
			},
		},
		{
			inA: []string{"a"},
			result: map[string]bool{
				"a": true,
			},
		},
		{
			inA:    []string{},
			result: map[string]bool{},
		},
	}
	for _, testCase := range testCases {
		result := ConvertStringSliceToMap(testCase.inA)
		if !reflect.DeepEqual(result, testCase.result) {
			t.Error(result, " should equal: ", testCase.result)
		}
	}

}
