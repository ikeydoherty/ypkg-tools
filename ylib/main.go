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
	"os/exec"
	"path"
	"regexp"
	"strings"
)

type SourceInfo struct {
	PkgName   string
	BaseName  string
	Version   string
	SourceURI string
}

const RootDirectory = "output_scan"

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

// Download the given file to the current directory
func FetchURI(source *SourceInfo) bool {
	cmd := exec.Command("curl", []string{"-o", source.BaseName, source.SourceURI, "--location"}...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return false
	}
	return true
}

// We're dealing with local files we create..
func PathExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

// Explode the tarball/zip/whathaveyou
func ExplodeSource(source *SourceInfo) (string, bool) {
	var cmd string

	// TODO: Use an absolute path, we need to use subdirs..
	tarball := source.BaseName
	outdir := "./" + RootDirectory

	// Ideally we need to nuke the old one.
	if !PathExists(outdir) {
		if err := os.MkdirAll(outdir, 00755); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create outdir: %v", err)
			return "", false
		}
	}

	if strings.HasSuffix(source.BaseName, ".zip") {
		cmd = fmt.Sprintf("unzip %v -d %s", tarball, outdir)
	} else if strings.Contains(source.BaseName, ".tar") {
		cmd = fmt.Sprintf("tar xf %v -C %s", tarball, outdir)
	} else {
		return "", false
	}

	coms := strings.Split(cmd, " ")
	command := exec.Command(coms[0], coms[1:]...)
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return "", false
	}

	return "", false
}
