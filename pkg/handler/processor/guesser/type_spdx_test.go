//
// Copyright 2022 The GUAC Authors.
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

package guesser

import (
	"testing"

	testdata "github.com/guacsec/guac/internal/testing/processor"
	"github.com/guacsec/guac/pkg/handler/processor"
)

func Test_spdxTypeGuesser_GuessDocumentType(t *testing.T) {
	testCases := []struct {
		name     string
		blob     []byte
		expected processor.DocumentType
	}{{
		name: "invalid spdx Document",
		blob: []byte(`{
			"abc": "def"
		}`),
		expected: processor.DocumentUnknown,
	}, {
		name:     "invalid spdx Document",
		blob:     testdata.SpdxInvalidExample,
		expected: processor.DocumentUnknown,
	}, {
		name:     "valid small spdx Document",
		blob:     testdata.SpdxExampleSmall,
		expected: processor.DocumentSPDX,
	}, {
		name:     "valid big spdx Document",
		blob:     testdata.SpdxExampleBig,
		expected: processor.DocumentSPDX,
	}}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			guesser := &spdxTypeGuesser{}
			f := guesser.GuessDocumentType(tt.blob, processor.FormatJSON)
			if f != tt.expected {
				t.Errorf("got the wrong format, got %v, expected %v", f, tt.expected)
			}
		})
	}
}
