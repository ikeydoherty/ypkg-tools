package ylib

import (
	"testing"
)

func TestExamineURI(t *testing.T) {

	uri_tests := make(map[string]SourceInfo)

	uri_tests["https://github.com/solus-project/linux-steam-integration/releases/download/v0.2/linux-steam-integration-0.2.tar.xz"] = SourceInfo{PkgName: "linux-steam-integration", Version: "0.2"}

	// github v URL
	uri_tests["https://github.com/solus-project/linux-steam-integration/archive/v0.2.tar.gz"] = SourceInfo{PkgName: "linux-steam-integration", Version: "0.2"}

	// pypi url with odd name package
	uri_tests["https://pypi.python.org/packages/fc/f1/7530ac8594453fc850e53580256f3152a8d8f2bb351bc3d0df8d7b53dbde/ruamel.yaml-0.11.11.tar.gz"] = SourceInfo{PkgName: "ruamel.yaml", Version: "0.11.11"}

	// Very odd sourceforge URL
	uri_tests["http://internode.dl.sourceforge.net/project/yodl/yodl/3.05.01/yodl_3.05.01.orig.tar.gz"] = SourceInfo{PkgName: "yodl", Version: "3.05.01.orig"}

	for uri, expected := range uri_tests {
		computed := ExamineURI(uri)
		if computed.PkgName != expected.PkgName {
			t.Fatalf("PkgName: Expected: %v, got: %v\n", expected.PkgName, computed.PkgName)
		}
		if computed.Version != expected.Version {
			t.Fatalf("Version: Expected: %v, got: %v\n", expected.Version, computed.Version)
		}
	}
}

func TestExamineSecond(t *testing.T) {
	t.Skip()
}
