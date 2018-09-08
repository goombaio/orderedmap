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

func BenchmarkOrderedMap_Put(b *testing.B) {
	m := orderedmap.NewOrderedMap()

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Put(i, i+1)
	}
}

var (
	resultBenchmarkOrderedMapGet1 interface{}
	resultBenchmarkOrderedMapGet2 bool
)

func BenchmarkOrderedMap_Get(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	for i := 0; i < b.N; i++ {
		m.Put(i, i+1)
	}

	var (
		value interface{}
		found bool
	)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Get(i)
	}
	b.StopTimer()

	resultBenchmarkOrderedMapGet1 = value
	resultBenchmarkOrderedMapGet2 = found
}

func BenchmarkOrderedMap_Remove(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	for i := 0; i < b.N; i++ {
		m.Put(i, i+1)
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Remove(i)
	}
}

var resultBenchmarkOrderedMapKeys []interface{}

func BenchmarkOrderedMap_Keys(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	for i := 0; i < b.N; i++ {
		m.Put(i, i+1)
	}

	var keys []interface{}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		keys = m.Keys()
	}
	b.StopTimer()

	resultBenchmarkOrderedMapKeys = keys
}

var resultBenchmarkOrderedMapValues []interface{}

func BenchmarkOrderedMap_Values(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	for i := 0; i < b.N; i++ {
		m.Put(i, i+1)
	}

	var values []interface{}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		values = m.Values()
	}
	b.StopTimer()

	resultBenchmarkOrderedMapValues = values
}

var resultBenchmarkOrderedMapSize int

func BenchmarkOrderedMap_Size(b *testing.B) {
	m := orderedmap.NewOrderedMap()
	for i := 0; i < b.N; i++ {
		m.Put(i, i+1)
	}

	var size int
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		size = m.Size()
	}
	b.StopTimer()
	resultBenchmarkOrderedMapSize = size
}
