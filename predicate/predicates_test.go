package predicate_test

// Licensed to BlueSoft Development, LLC under one or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information regarding copyright ownership.  BlueSoft Development, LLC
// licenses this file to you under the Apache License, Version 2.0 (the "License"); you may not use this file except in
// compliance with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations under the License.

import (
	. "go-http-matchers/predicate"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestTrue(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	assert.True(t, True().Accept(req), "expected true.")
}

func TestFalse(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	assert.False(t, False().Accept(req), "expected false.")
}

func TestNot(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	assert.False(t, Not(True()).Accept(req))
	assert.True(t, Not(False()).Accept(req))
}

func TestAnd(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	assert.True(t, And(True(), True()).Accept(req), "expected true.")
	assert.True(t, And(True(), True(), True()).Accept(req), "expected false.")
	assert.False(t, And(False(), True()).Accept(req), "expected false.")
	assert.False(t, And(True(), False()).Accept(req), "expected false.")
	assert.False(t, And(True(), True(), False()).Accept(req), "expected false.")
	assert.False(t, And(True(), False(), True()).Accept(req), "expected false.")
	assert.False(t, And(False(), True(), True()).Accept(req), "expected false.")
}

func TestOr(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	assert.True(t, Or(True(), True()).Accept(req), "expected true.")
	assert.True(t, Or(True(), True(), True()).Accept(req), "expected true.")
	assert.True(t, Or(False(), True()).Accept(req), "expected true.")
	assert.True(t, Or(True(), False()).Accept(req), "expected true.")
	assert.True(t, Or(True(), True(), False()).Accept(req), "expected true.")
	assert.True(t, Or(True(), False(), True()).Accept(req), "expected true.")
	assert.True(t, Or(False(), True(), True()).Accept(req), "expected true.")
	assert.False(t, Or(False(), False()).Accept(req), "expected false.")
	assert.False(t, Or(False(), False(), False()).Accept(req), "expected false.")
}

func TestMethodIs(t *testing.T) {
	req, err := http.NewRequest("GET", "http://foo.com/test/foo/bar?q=5&l=3", nil)
	assert.NoError(t, err, "failed to create test request.")
	assert.True(t, MethodIs("GET").Accept(req))
	assert.False(t, MethodIs("POST").Accept(req))
}

