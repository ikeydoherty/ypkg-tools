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
	if source == nil {
		os.Exit(0)
	}
	if ylib.PathExists(source.BaseName) {
		if err := os.Remove(source.BaseName); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to unlink %v: %s", source.SourceURI, err)
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

	if len(args) < 2 {
		usage_and_quit(args)
	}

	url := args[1]
	source_info := ylib.ExamineURI(url)
	defer cleanup_and_exit(source_info)
	if source_info == nil {
		fmt.Fprintf(os.Stderr, "Failed to examine %v\n", url)
		return
	}

	if !ylib.FetchURI(source_info) {
		return
	}

	badness = false
}
