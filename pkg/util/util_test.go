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
	"strconv"
	"testing"
)

func TestUnionMap(t *testing.T) {
	a := map[string]bool{
		"a": true,
		"b": true,
	}
	b := map[string]bool{
		"c": true,
	}
	UnionMap(a, b)
	results := []string{"a", "b", "c"}

	for _, result := range results {
		if _, ok := a[result]; !ok {
			t.Error("unionMap should contain: " + result)
			t.FailNow()
		}
	}

}
func TestSubtractMap(t *testing.T) {
	a := map[string]bool{
		"a": true,
		"b": true,
	}
	b := map[string]bool{
		"b": true,
	}
	result := SubtractMap(a, b)

	if len(result) != 1 {
		t.Error("SubtractMap should contain only one string. It contains: " + strconv.Itoa(len(result)))
		t.FailNow()
	}

	if _, ok := result["a"]; !ok {
		t.Error("unionMap should contain 'a' string")
		t.FailNow()
	}

}
