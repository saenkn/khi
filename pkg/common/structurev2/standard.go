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

// standard.go contains types to hold structured data on memory as its field and implements Node interface.

type StandardScalarNode[T any] struct {
	value T
}

func (n *StandardScalarNode[T]) Type() NodeType {
	return ScalarNodeType
}

func (n *StandardScalarNode[T]) NodeScalarValue() (any, error) {
	return n.value, nil
}

func (n *StandardScalarNode[T]) Children() NodeChildrenIterator {
	return func(func(key NodeChildrenKey, value Node) bool) {}
}

type StandardSequenceNode struct {
	value []Node
}

func (n *StandardSequenceNode) Type() NodeType {
	return SequenceNodeType
}

func (n *StandardSequenceNode) NodeScalarValue() (any, error) {
	return nil, ErrNonScalarNode
}

func (n *StandardSequenceNode) Children() NodeChildrenIterator {
	return func(f func(key NodeChildrenKey, value Node) bool) {
		for i, v := range n.value {
			if !f(NodeChildrenKey{Index: i}, v) {
				return
			}
		}
	}
}

type StandardMapNode struct {
	// keys is the list of keys in values.
	// Needed because the key order of map is not assured in Go.
	keys   []string
	values []Node
}

func (n *StandardMapNode) Type() NodeType {
	return MapNodeType
}

func (n *StandardMapNode) NodeScalarValue() (any, error) {
	return nil, ErrNonScalarNode
}

func (n *StandardMapNode) Children() NodeChildrenIterator {
	return func(f func(key NodeChildrenKey, value Node) bool) {
		for i, k := range n.keys {
			if !f(NodeChildrenKey{Index: i, Key: k}, n.values[i]) {
				return
			}
		}
	}
}
