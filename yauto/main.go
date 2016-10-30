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
	"strings"
)

func main() {
	fmt.Fprintf(os.Stderr, "Not yet implemented\n")

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [primary archive] [additional]\n", os.Args[0])
		os.Exit(1)
	}

	// All archives
	archives := os.Args[1:]
	primary := archives[0]

	context, err := ylib.NewContext(".")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot create context: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("WorkDir: %s\n", context.WorkingDirectory)

	fmt.Fprintf(os.Stderr, "Primary archive: %s\n", primary)
	fmt.Fprintf(os.Stderr, "Fetching: %s\n", strings.Join(archives, ", "))

	c := ylib.MarshalledYpkg{Name: "i-r-package", Version: "1.2.2"}
	c.Description = "Herp\nDerp"
	c.Summary = "I am a summary.\nThis is my second\nand third line"
	c.Source = make(map[string]string)
	c.Source["http://example.com"] = "SoMeHaSh"
	if err := ylib.WriteYpkg("package.yml", &c); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	os.Exit(1)
}
