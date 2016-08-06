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

func usage_and_quit(args []string) {
	fmt.Printf("Usage: %v [url]\n", args[0])
	os.Exit(1)
}

func main() {
	args := os.Args

	if len(args) < 2 {
		usage_and_quit(args)
	}

	// url := args[1]
	url := "https://github.com/solus-project/linux-steam-integration/releases/download/v0.2/linux-steam-integration-0.2.tar.xz"
	// url := "https://github.com/solus-project/linux-steam-integration/archive/v0.2.tar.gz"
	// url := "https://pypi.python.org/packages/fc/f1/7530ac8594453fc850e53580256f3152a8d8f2bb351bc3d0df8d7b53dbde/ruamel.yaml-0.11.11.tar.gz"
	//url := "http://internode.dl.sourceforge.net/project/yodl/yodl/3.05.01/yodl_3.05.01.orig.tar.gz"
	source_info := ylib.ExamineURI(url)
	if source_info == nil {
		os.Exit(1)
	}
	fmt.Printf("Got somethin: %v (%v): %v\n", source_info.PkgName, source_info.Version, source_info.BaseName)
}
