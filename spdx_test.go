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
	"regexp"
	"slices"
	"strings"
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

func TestLicenseForID_Deprecated(t *testing.T) {
	t.Parallel()

	got, ok := LicenseForID("gpl-3.0")
	if !ok {
		t.Fatalf("failed to get license info")
	}
	if !got.Deprecated {
		t.Fatalf("expected license %s to be deprecated", got.ID)
	}
}

func TestIDsMatchExpectedRegext(t *testing.T) {
	t.Parallel()

	regexp, err := regexp.Compile("^[a-zA-Z0-9-.+]+$")
	if err != nil {
		t.Fatal(err.Error())
	}
	for _, license := range AllLicenses() {
		if !regexp.Match([]byte(license.ID)) {
			t.Fatalf("license ID %q did not match regex", license.ID)
		}
	}
}

func TestLicenseForID_CaseInsensitive(t *testing.T) {
	t.Parallel()

	for _, license := range AllLicenses() {
		for _, id := range []string{license.ID, strings.ToLower(license.ID), strings.ToUpper(license.ID)} {
			got, ok := LicenseForID(id)
			if !ok {
				t.Fatalf("failed to get license info for %q", id)
			}
			if !reflect.DeepEqual(got, license) {
				t.Fatalf("got %+v for %q, want %+v", got, id, license)
			}
		}
	}
}

func TestLicenseForID_NotFound(t *testing.T) {
	t.Parallel()

	for _, id := range []string{"", "not-a-license", "apache-2.0x", "apache-2", " MIT", "zzzz", "Apache\r2.0"} {
		if license, ok := LicenseForID(id); ok {
			t.Fatalf("expected no license for %q, got %+v", id, license)
		}
	}
}

func TestAllLicenses_SortedByID(t *testing.T) {
	t.Parallel()

	licenses := AllLicenses()
	if len(licenses) == 0 {
		t.Fatal("expected licenses")
	}
	if !slices.IsSortedFunc(licenses, func(a License, b License) int {
		return strings.Compare(a.ID, b.ID)
	}) {
		t.Fatal("expected licenses to be sorted by ID")
	}
}

func TestLicenseEntriesSortedByLowercaseID(t *testing.T) {
	t.Parallel()

	for i := 1; i < len(licenseEntriesByLowercaseID); i++ {
		previous, current := licenseEntriesByLowercaseID[i-1].id, licenseEntriesByLowercaseID[i].id
		if strings.ToLower(previous) >= strings.ToLower(current) {
			t.Fatalf("licenseEntriesByLowercaseID not strictly sorted by lowercase ID: %q, %q", previous, current)
		}
	}
}

func TestCompareASCIIFold(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		a    string
		b    string
		want int
	}{
		{"", "", 0},
		{"MIT", "mit", 0},
		{"a", "B", -1},
		{"B", "a", 1},
		{"abc", "ABCD", -1},
		{"ABCD", "abc", 1},
		{"Zlib", "zlib-acknowledgement", -1},
		{"0BSD", "aal", -1},
	}
	for _, testCase := range testCases {
		if got := compareASCIIFold(testCase.a, testCase.b); got != testCase.want {
			t.Errorf("compareASCIIFold(%q, %q) = %d, want %d", testCase.a, testCase.b, got, testCase.want)
		}
		want := strings.Compare(strings.ToLower(testCase.a), strings.ToLower(testCase.b))
		if got := compareASCIIFold(testCase.a, testCase.b); got != want {
			t.Errorf("compareASCIIFold(%q, %q) = %d, strings.Compare on lowercase = %d", testCase.a, testCase.b, got, want)
		}
	}
}
