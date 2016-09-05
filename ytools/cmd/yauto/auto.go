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

package yauto

import (
	"flag"
	"github.com/ikeydoherty/ypkg-tools/ytools/cmd"
)

var set *flag.FlagSet

func usage() {
	println("Usage: ytools yauto <URI>")
	set.PrintDefaults()
}

func flags() *flag.FlagSet {
	set = flag.NewFlagSet("auto", flag.ContinueOnError)
	set.Usage = usage
	return set
}

func run() {
	println("Ran auto")
}

// GetCmd - Builds the "auto" command
func GetCmd() *cmd.CMD {
	c := &cmd.CMD{}
	c.Flags = flags()
	c.Run = run
	return c
}
