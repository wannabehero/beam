// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package primitives

import (
	"testing"

	"github.com/apache/beam/sdks/v2/go/pkg/beam/testing/ptest"
	"github.com/apache/beam/sdks/v2/go/test/integration"
)

func TestValueState(t *testing.T) {
	integration.CheckFilters(t)
	ptest.RunAndValidate(t, ValueStateParDo())
}

func TestValueState_Clear(t *testing.T) {
	integration.CheckFilters(t)
	ptest.RunAndValidate(t, ValueStateParDo_Clear())
}

func TestBagState(t *testing.T) {
	integration.CheckFilters(t)
	ptest.RunAndValidate(t, BagStateParDo())
}

func TestBagState_Clear(t *testing.T) {
	integration.CheckFilters(t)
	ptest.RunAndValidate(t, BagStateParDo_Clear())
}

func TestCombiningState(t *testing.T) {
	integration.CheckFilters(t)
	ptest.RunAndValidate(t, CombiningStateParDo())
}

func TestMapState(t *testing.T) {
	integration.CheckFilters(t)
	ptest.RunAndValidate(t, MapStateParDo())
}
