// Copyright 2015 Peter Goetz
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

package mockgen_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	mockgen "github.com/petergtz/pegomock/mockgen"
)

func TestMockGeneration(t *testing.T) {
	RunSpecs(t, "Mock generation suite")
}

var _ = It("Generate mocks", func() {
	mockgen.Run("",
		"../pegomock/mock_display_test.go", "pegomock_test",
		"",
		false,
		"github.com/petergtz/pegomock/test_interface", "Display")
})
