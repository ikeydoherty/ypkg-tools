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
	"strings"
	"sync"
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
	re := regexp.MustCompile(`https://github.com/.*/(.*?)/archive/v?(.*).(tar|zip)`)
	if ret := re.FindStringSubmatch(uri); len(ret) > 0 {
		return &SourceInfo{SourceURI: uri, BaseName: basename, PkgName: ret[1], Version: ret[2]}
	}

	// Gitlab, especially special.
	re = regexp.MustCompile(`https://gitlab.com/.*/(.*?)/repository/archive.[tar|zip].*\?ref=v?(.*)`)
	if ret := re.FindStringSubmatch(uri); len(ret) > 0 {
		return &SourceInfo{SourceURI: uri, BaseName: basename, PkgName: ret[1], Version: ret[2]}
	}

	// Try a "normal path"
	re = regexp.MustCompile(`([a-zA-Z-_.0-9]+)[-|_](.*?)\.(t|zip).*`)
	if ret := re.FindStringSubmatch(basename); len(ret) > 0 {
		return &SourceInfo{SourceURI: uri, BaseName: basename, PkgName: ret[1], Version: ret[2]}
	}


	return nil
}

// Actual goroutine to scan the files
func scan_path(path string, info os.FileInfo, wg *sync.WaitGroup) {
	defer wg.Done()

	lpath := strings.ToLower(info.Name())

	if strings.HasPrefix(lpath, "license") || strings.HasPrefix(lpath, "licence") || strings.HasPrefix(lpath, "copying") {
		fmt.Printf("License encountered: %s\n", path)
	}
}

// Scan the tree to find things of interest.
// At some point we need to return the results or do something
// useful with them.
func ScanTree(rootdir string) bool {
	var wg sync.WaitGroup

	// Need access to the waitgroup ^
	wfunc := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Don't want to scan directories.
		if info.IsDir() {
			return nil
		}
		wg.Add(1)
		go scan_path(path, info, &wg)
		return nil
	}

	if err := filepath.Walk(rootdir, wfunc); err != nil {
		fmt.Fprintf(os.Stderr, "Hit an error. %v", err)
	}

	wg.Wait()
	return false
}
