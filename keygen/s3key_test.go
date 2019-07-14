// Copyright 2019 h2cone https://github.com/h2cone

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package keygen

import (
	"path/filepath"
	"testing"
)

func TestUUIDWithExt(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"golang.png", ".png"},
		{"golang", ""},
		{"golang.png.jpg", ".jpg"},
	}
	for _, test := range tests {
		if got := filepath.Ext(UUIDWithExt(test.input)); got != test.want {
			t.Errorf("UUIDWithExt(%s) = %s", test.input, got)
		}
	}
}
