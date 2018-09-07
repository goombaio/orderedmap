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

package orderedmap

import (
	"sync"
)

// OrderedMap insertion ordered Map implementation
type OrderedMap struct {
	sync.Mutex

	store map[interface{}]interface{}
	keys  []interface{}
}

// NewOrderedMap return a new Map implemented by OrderedMap
func NewOrderedMap() *OrderedMap {
	m := &OrderedMap{
		store: make(map[interface{}]interface{}),
		keys:  make([]interface{}, 0, 0),
	}

	return m
}

// Put add a key-value pair to the OrderedMap
func (m *OrderedMap) Put(key interface{}, value interface{}) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.store[key]; !ok {
		m.keys = append(m.keys, key)
	}

	m.store[key] = value
}

// Get return the value of a key from the OrderedMap
func (m *OrderedMap) Get(key interface{}) (value interface{}, found bool) {
	m.Lock()
	defer m.Unlock()

	value, found = m.store[key]
	return value, found
}

// Remove remove a key-value pair from the OrderedMap
func (m *OrderedMap) Remove(key interface{}) {
	m.Lock()
	defer m.Unlock()

	if _, found := m.store[key]; !found {
		return
	}

	delete(m.store, key)

	for i := range m.keys {
		if m.keys[i] == key {
			m.keys = append(m.keys[:i], m.keys[i+1:]...)
			break
		}
	}
}

// Empty return if the OrderedMap in empty or not
func (m *OrderedMap) Empty() bool {
	m.Lock()
	defer m.Unlock()

	return len(m.store) == 0
}

// Keys return the keys of the OrderedMap in insertion order
func (m *OrderedMap) Keys() []interface{} {
	m.Lock()
	defer m.Unlock()

	return m.keys
}

// Values return the values of the OrderedMap in insertion order
func (m *OrderedMap) Values() []interface{} {
	m.Lock()
	defer m.Unlock()

	values := make([]interface{}, len(m.store))
	for i, key := range m.keys {
		values[i] = m.store[key]
	}
	return values
}

// Size return the size of the OrderedMap
func (m *OrderedMap) Size() int {
	m.Lock()
	defer m.Unlock()

	return len(m.store)
}
