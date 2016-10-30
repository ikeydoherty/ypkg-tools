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

package main

import (
	"fmt"
	"github.com/ikeydoherty/ypkg-tools/ylib"
	"os"
)

// Just to track inside the exit
var badness bool

func usageAndQuit(args []string) {
	fmt.Printf("Usage: %v [url]\n", args[0])
	os.Exit(1)
}

func cleanupAndExit(source *ylib.SourceInfo) {
	// Remove the explode tree if it exists
	if ylib.PathExists("./" + ylib.RootDirectory) {
		if err := os.RemoveAll("./" + ylib.RootDirectory); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to nuke tree: %v\n", err)
		}
	}
	if source == nil {
		os.Exit(0)
	}
	if ylib.PathExists(source.BaseName) {
		if err := os.Remove(source.BaseName); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to unlink %v: %s\n", source.SourceURI, err)
			os.Exit(1)
		}
	}
	if badness {
		os.Exit(1)
	}
}

func main() {
	args := os.Args
	badness = true

	// i.e. yauto $URL
	if len(args) < 2 {
		usageAndQuit(args)
	}

	// Try and inspect the URL now, rather than later on..
	baseURL := args[1]
	url, err := ylib.StripURI(baseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to strip URL fragments: %v\n", err)
		return
	}

	sourceInfo := ylib.ExamineURI(url)
	defer cleanupAndExit(sourceInfo)
	// No idea how to read this
	if sourceInfo == nil {
		fmt.Fprintf(os.Stderr, "Failed to examine %v\n", url)
		return
	}

	// Try and download the source
	if !ylib.FetchURI(sourceInfo) {
		return
	}

	rootdir, success := ylib.ExplodeSource(sourceInfo)
	if !success {
		fmt.Fprintf(os.Stderr, "Failed to explode source\n")
		return
	}

	if !ylib.ScanTree(rootdir) {
		fmt.Fprintf(os.Stderr, "Failed to scan tree\n")
		return
	}
	fmt.Println("Not fully implemented")

	badness = false
}
