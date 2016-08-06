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
	"bufio"
	"fmt"
	"os"
)

// License mapping table for derp detection
var license_table map[string]string

// Hash mapping for known SPDX licenses
var license_hash map[string]string

func init() {
	license_table = make(map[string]string)
	license_hash = make(map[string]string)
}

// Scan a file for license text
func scan_license(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		return ""
	}
	defer fi.Close()

	sc := bufio.NewScanner(fi)
	for sc.Scan() {
		line := sc.Text()
		fmt.Println(line)
	}
	return ""
}
