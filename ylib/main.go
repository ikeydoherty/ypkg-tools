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
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

type SourceInfo struct {
	PkgName   string
	BaseName  string
	Version   string
	SourceURI string
}

const RootDirectory = "output_scan"

// Examine the URI and try to learn the valid version and name for this Thing
func ExamineURI(uri string) *SourceInfo {
	basename := path.Base(uri)

	// Try github v* match
	re := regexp.MustCompile(`https://github.com/.*/(.*?)/archive/v?(.*).tar`)
	if ret := re.FindStringSubmatch(uri); len(ret) > 0 {
		return &SourceInfo{SourceURI: uri, BaseName: basename, PkgName: ret[1], Version: ret[2]}
	}

	// Try a "normal path"
	re = regexp.MustCompile(`([a-zA-Z-_.0-9]+)[-|_](.*?)\.[t|zip].*`)
	if ret := re.FindStringSubmatch(basename); len(ret) > 0 {
		return &SourceInfo{SourceURI: uri, BaseName: basename, PkgName: ret[1], Version: ret[2]}
	}

	return nil
}

// Scan the tree to find things of interest.
// At some point we need to return the results or do something
// useful with them.
func ScanTree(rootdir string) bool {
	wfunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		return nil
	}

	if err := filepath.Walk(rootdir, wfunc); err != nil {
		fmt.Fprintf(os.Stderr, "Hit an error. %v", err)
	}
	return false
}
