// Copyright 2018, Goomba project Authors. All rights reserved.
//
// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with this
// work for additional information regarding copyright ownership.  The ASF
// licenses this file to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the
// License for the specific language governing permissions and limitations
// under the License.

package orderedmap_test

import (
	"testing"

	"github.com/goombaio/orderedmap"
)

type customType struct {
	foo string
}

func TestOrderedMap_Put(t *testing.T) {
	m := orderedmap.NewOrderedMap()

	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")

	m.Put(1, "a") //overwrite
	m.Put(2, "b")

	structKey := customType{"skey"}
	structValue := customType{"svalue"}
	m.Put(structKey, structValue)
	m.Put(&structKey, &structValue)

	m.Put(true, false)
}

func TestOrderedMap_Put_overwrite(t *testing.T) {
	m := orderedmap.NewOrderedMap()

	m.Put(1, "x")
	m.Put(1, "a") //overwrite

	actualValue, actualFound := m.Get(1)
	if actualValue != "a" || !actualFound {
		t.Errorf("Got %v expected %v", actualValue, true)
	}
}

func TestOrderedMap_Get(t *testing.T) {
	m := orderedmap.NewOrderedMap()

	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")

	m.Put(1, "a") //overwrite
	m.Put(2, "b")

	structKey := customType{"skey"}
	structValue := customType{"svalue"}
	m.Put(structKey, structValue)
	m.Put(&structKey, &structValue)

	m.Put(true, false)

	table := []struct {
		key           interface{}
		expectedValue interface{}
		expectedFound bool
	}{
		{1, "a", true},
		{2, "b", true},
		{3, "c", true},
		{4, "d", true},
		{5, "e", true},
		{6, "f", true},
		{7, "g", true},
		{8, nil, false},
		{structKey, structValue, true},
		{&structKey, &structValue, true},
		{true, false, true},
	}

	for _, test := range table {
		actualValue, actualFound := m.Get(test.key)
		if actualValue != test.expectedValue || actualFound != test.expectedFound {
			t.Errorf("Got %v expected %v", actualValue, test.expectedValue)
		}
	}
}

func TestOrderedMap_Remove(t *testing.T) {
	m := orderedmap.NewOrderedMap()
	m.Put("bar", "foo")
	m.Put("foo", "bar")

	actualValue, actualFound := m.Get("foo")
	if actualValue != "bar" || !actualFound {
		t.Errorf("Got %v:%v expected %v:%v", actualValue, actualFound, "bar", true)
	}

	m.Remove("foo")
	actualValue, actualFound = m.Get("foo")
	if actualValue != nil || actualFound {
		t.Errorf("Got %v:%v expected %v:%v", actualValue, actualFound, nil, false)
	}

	m.Remove("foo") // already removed
	actualValue, actualFound = m.Get("foo")
	if actualValue != nil || actualFound {
		t.Errorf("Got %v:%v expected %v:%v", actualValue, actualFound, nil, false)
	}
}

func TestOrderedMap_Empty(t *testing.T) {
	m := orderedmap.NewOrderedMap()
	actualValue := m.Empty()
	if !actualValue {
		t.Errorf("Got %v expected %v", actualValue, true)
	}

	m.Put("foo", "bar")
	actualValue = m.Empty()
	if actualValue {
		t.Errorf("Got %v expected %v", actualValue, false)
	}
}

func TestOrderedMap_Size(t *testing.T) {
	m := orderedmap.NewOrderedMap()

	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")

	m.Put(1, "a") //overwrite
	m.Put(2, "b")

	structKey := customType{"skey"}
	structValue := customType{"svalue"}
	m.Put(structKey, structValue)
	m.Put(&structKey, &structValue)

	m.Put(true, false)

	if actualSize := m.Size(); actualSize != 10 {
		t.Errorf("Got %v expected %v", actualSize, 10)
	}
}

func TestOrderedMap_Keys(t *testing.T) {
	m := orderedmap.NewOrderedMap()

	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")

	m.Put(1, "a") //overwrite
	m.Put(2, "b")

	structKey := customType{"skey"}
	structValue := customType{"svalue"}
	m.Put(structKey, structValue)
	m.Put(&structKey, &structValue)

	m.Put(true, false)

	actualKeys := m.Keys()
	expectedKeys := []interface{}{5, 6, 7, 3, 4, 1, 2, structKey, &structKey, true}
	if !sameElements(actualKeys, expectedKeys) {
		t.Errorf("Got %v expected %v", actualKeys, expectedKeys)
	}
}

func TestOrderedMap_Values(t *testing.T) {
	m := orderedmap.NewOrderedMap()

	m.Put(5, "e")
	m.Put(6, "f")
	m.Put(7, "g")
	m.Put(3, "c")
	m.Put(4, "d")
	m.Put(1, "x")
	m.Put(2, "b")

	m.Put(1, "a") //overwrite
	m.Put(2, "b")

	structKey := customType{"skey"}
	structValue := customType{"svalue"}
	m.Put(structKey, structValue)
	m.Put(&structKey, &structValue)

	m.Put(true, false)

	actualValues := m.Values()
	expectedValues := []interface{}{"e", "f", "g", "c", "d", "a", "b", structValue, &structValue, false}
	if !sameElements(actualValues, expectedValues) {
		t.Errorf("Got %v expected %v", actualValues, expectedValues)
	}
}

func TestOrderedMap_String(t *testing.T) {
	m := orderedmap.NewOrderedMap()
	m.Put(1, "foo")
	m.Put(2, "bar")

	expected := "[1:foo 2:bar]"
	result := m.String()
	if expected != result {
		t.Fatalf("Got %q expected %q", result, expected)
	}
}

func sameElements(a []interface{}, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for _, av := range a {
		found := false
		for _, bv := range b {
			if av == bv {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
