// Copyright © 2020 Banzai Cloud
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

package api

import (
	"net/http"
	"path/filepath"
	"strings"

	"github.com/markbates/pkger"
)

type pkgerFileSystem struct {
	dir pkger.Dir
}

func (fs pkgerFileSystem) Open(name string) (http.File, error) {
	return fs.dir.Open(name)
}

func (fs pkgerFileSystem) Exists(prefix string, filePath string) bool {
	if p := strings.TrimPrefix(filePath, prefix); len(p) < len(filePath) {
		if p == "index.html" {
			return false
		}

		stat, err := pkger.Stat(filepath.Join(string(fs.dir), p))
		if err != nil {
			return false
		}

		return !stat.IsDir()
	}

	return false
}
