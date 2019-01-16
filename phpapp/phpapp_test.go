/*
 * Copyright 2018-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package phpapp

import (
	"testing"

	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestUnitPHPApp(t *testing.T) {
	RegisterTestingT(t)
	spec.Run(t, "PHPApp", testPHPApp, spec.Report(report.Terminal{}))
}

func testPHPApp(t *testing.T, when spec.G, it spec.S) {
	when("a version is set", func() {
		it("uses php.version from buildpack.yml, if set", func() {
		})
	})
}
