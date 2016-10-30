//
// Copyright © 2016 Ikey Doherty
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
//

package ylib

import (
	"path/filepath"
)

// Context provides ways to ensure absolute paths formed from the base WorkingDirectory
type Context struct {
	WorkingDirectory string
}

// NewContext returns a new operating context for library consumers
func NewContext(base string) (*Context, error) {
	// Assume current directory
	if base == "" {
		base = "."
	}

	absPath, err := filepath.Abs(base)
	if err != nil {
		return nil, err
	}

	return &Context{WorkingDirectory: absPath}, nil
}

// GetExtractionRoot returns the directory used for archive extraction and examination
func (c *Context) GetExtractionRoot() string {
	return filepath.Join(c.WorkingDirectory, "_extract_root")
}
