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

func usage_and_quit(args []string) {
	fmt.Printf("Usage: %v [url]\n", args[0])
	os.Exit(1)
}

func cleanup_and_exit(source *ylib.SourceInfo) {
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
		usage_and_quit(args)
	}

	// Try and inspect the URL now, rather than later on..
	url := args[1]
	source_info := ylib.ExamineURI(url)
	defer cleanup_and_exit(source_info)
	// No idea how to read this
	if source_info == nil {
		fmt.Fprintf(os.Stderr, "Failed to examine %v\n", url)
		return
	}

	// Try and download the source
	if !ylib.FetchURI(source_info) {
		return
	}

	rootdir, success := ylib.ExplodeSource(source_info)
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
