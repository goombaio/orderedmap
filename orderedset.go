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

package orderedset

import (
	"sync"

	"github.com/goombaio/orderedmap"
)

// OrderedSet insertion ordered Set implementation
type OrderedSet struct {
	sync.Mutex

	// the underlying store of the Set
	store *orderedmap.OrderedMap
	index map[interface{}]int
	// currentIndex keeps track of the keys of the underlying store
	currentIndex int
}

// NewOrderedSet return a new Set implemented by OrderedSet
func NewOrderedSet() *OrderedSet {
	return &OrderedSet{
		store: orderedmap.NewOrderedMap(),
		index: make(map[interface{}]int),
	}
}

// Add add items to the OrderedSet
func (s *OrderedSet) Add(items ...interface{}) {
	for _, item := range items {
		if _, found := s.index[item]; found {
			continue
		}

		s.put(item)
	}
}

// Remove remove items from the OrderedSet
func (s *OrderedSet) Remove(items ...interface{}) {
	for _, item := range items {
		index, found := s.index[item]
		if !found {
			return
		}

		s.remove(item, index)
	}
}

// Contains return if OrderedSet contains the specified items or not
func (s *OrderedSet) Contains(items ...interface{}) bool {
	for _, item := range items {
		if _, found := s.index[item]; !found {
			return false
		}
	}
	return true
}

// Empty return if the OrderedSet in empty or not
func (s *OrderedSet) Empty() bool {
	return s.store.Empty()
}

// Values return the values of the OrderedSet in insertion order
func (s *OrderedSet) Values() []interface{} {
	return s.store.Values()
}

// Size return the size of the OrderedSet
func (s *OrderedSet) Size() int {
	return s.store.Size()
}

func (s *OrderedSet) put(item interface{}) {
	s.Lock()
	defer s.Unlock()

	s.store.Put(s.currentIndex, item)
	s.index[item] = s.currentIndex
	s.currentIndex++
}

func (s *OrderedSet) remove(item interface{}, index int) {
	s.Lock()
	defer s.Unlock()

	s.store.Remove(index)
	delete(s.index, item)
}
