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
	"github.com/google/uuid"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// UUIDWithExt UUID combined file extension
func UUIDWithExt(filename string) string {
	id, err := uuid.NewRandom()
	if err != nil {
		return filename
	}
	ext := filepath.Ext(filename)
	return strings.ReplaceAll(id.String(), "-", "") + ext
}

// ContentMD5WithUnixNano .
func ContentMD5WithUnixNano(filename string) string {
	// TODO: ContentMD5_Nanosecond
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}
