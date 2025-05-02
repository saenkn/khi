// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package structurev2

import (
	"testing"

	_ "github.com/GoogleCloudPlatform/khi/internal/testflags"
)

func TestStandardSequenceNodeChildren(t *testing.T) {
	input := []string{"a", "b", "c"}
	node := StandardSequenceNode{value: []Node{
		&StandardScalarNode[string]{value: "a"},
		&StandardScalarNode[string]{value: "b"},
		&StandardScalarNode[string]{value: "c"},
	}}
	for key, value := range node.Children() {
		childValue, err := value.NodeScalarValue()
		if err != nil {
			t.Error(err)
		}

		if childValue != input[key.Index] {
			t.Errorf("expected %s, got %s", input[key.Index], childValue)
		}
	}
}

func TestStandardMappingNodeChildren(t *testing.T) {
	input := map[string]int{
		"b": 1,
		"a": 2,
		"c": 3,
	}
	indices := map[string]int{
		"b": 0,
		"a": 1,
		"c": 2,
	}
	node := StandardMapNode{values: []Node{
		&StandardScalarNode[int]{value: 1},
		&StandardScalarNode[int]{value: 2},
		&StandardScalarNode[int]{value: 3},
	}, keys: []string{"b", "a", "c"}}
	for key, value := range node.Children() {
		childValueAny, err := value.NodeScalarValue()
		if err != nil {
			t.Error(err)
		}
		if childValueInt, ok := childValueAny.(int); !ok {
			t.Errorf("expected int, got %T", childValueAny)
		} else if childValueInt != input[key.Key] {
			t.Errorf("expected %d, got %d", input[key.Key], childValueInt)
		}

		if key.Index != indices[key.Key] {
			t.Errorf("expected %d, got %d", indices[key.Key], key.Index)
		}
	}
}
