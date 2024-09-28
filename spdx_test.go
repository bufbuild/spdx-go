// Copyright 2024 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package spdx

import (
	"reflect"
	"testing"
)

func TestLicenseForID(t *testing.T) {
	t.Parallel()

	got, ok := LicenseForID("apache-2.0")
	if !ok {
		t.Fatalf("failed to get license info")
	}
	want := License{
		ID:              "Apache-2.0",
		Name:            "Apache License 2.0",
		Reference:       "https://spdx.org/licenses/Apache-2.0.html",
		ReferenceNumber: 501,
		DetailsURL:      "https://spdx.org/licenses/Apache-2.0.json",
		Deprecated:      false,
		SeeAlso:         []string{"https://www.apache.org/licenses/LICENSE-2.0", "https://opensource.org/licenses/Apache-2.0"},
		OSIApproved:     true,
	}
	if !reflect.DeepEqual(got, want) {
		t.Logf("got: %+v", got)
		t.Logf("want: %+v", want)
		t.Fatal("failed to get expected license info")
	}
}
