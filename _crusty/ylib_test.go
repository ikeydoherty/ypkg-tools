package ylib

import (
	"testing"
)

func TestExamineURI(t *testing.T) {

	uriTests := make(map[string]SourceInfo)

	uriTests["https://github.com/solus-project/linux-steam-integration/releases/download/v0.2/linux-steam-integration-0.2.tar.xz"] = SourceInfo{PkgName: "linux-steam-integration", Version: "0.2"}

	// github v URL tarball
	uriTests["https://github.com/solus-project/linux-steam-integration/archive/v0.2.tar.gz"] = SourceInfo{PkgName: "linux-steam-integration", Version: "0.2"}

	// github v URL zip
	uriTests["https://github.com/solus-project/linux-steam-integration/archive/v0.2.zip"] = SourceInfo{PkgName: "linux-steam-integration", Version: "0.2"}

	// pypi url with odd name package
	uriTests["https://pypi.python.org/packages/fc/f1/7530ac8594453fc850e53580256f3152a8d8f2bb351bc3d0df8d7b53dbde/ruamel.yaml-0.11.11.tar.gz"] = SourceInfo{PkgName: "ruamel.yaml", Version: "0.11.11"}

	// Very odd sourceforge URL
	uriTests["http://internode.dl.sourceforge.net/project/yodl/yodl/3.05.01/yodl_3.05.01.orig.tar.gz"] = SourceInfo{PkgName: "yodl", Version: "3.05.01"}

	// gitlab tarball
	uriTests["https://gitlab.com/manaplus/manaplus/repository/archive.tar.gz?ref=v1.6.7.30"] = SourceInfo{PkgName: "manaplus", Version: "1.6.7.30"}

	// gitbal zip
	uriTests["https://gitlab.com/manaplus/manaplus/repository/archive.zip?ref=v1.6.7.30"] = SourceInfo{PkgName: "manaplus", Version: "1.6.7.30"}

	// Mangled pypgi with md5sum
	uriTests["https://pypi.python.org/packages/ec/51/df65c73f957375bbd814bb0353213d0422d07fe538b92a2a23fa683e499b/meson-0.33.0.tar.gz#md5=3252395e6df14e6f85270abb2542e49b"] = SourceInfo{PkgName: "meson", Version: "0.33.0"}

	// dodgier sourceforge
	uriTests["http://netix.dl.sourceforge.net/project/ufoai/UFO_AI%202.x/2.5/ufoai-2.5-source.tar.bz2"] = SourceInfo{PkgName: "ufoai", Version: "2.5"}

	uriTests["https://bitbucket.org/sinbad/ogre/get/v1-9-0.tar.gz"] = SourceInfo{PkgName: "ogre", Version: "1.9.0"}

	for uri, expected := range uriTests {
		computed := ExamineURI(uri)
		if computed == nil {
			t.Fatalf("Failed to parse %v\n", uri)
		}
		if computed.PkgName != expected.PkgName {
			//t.Fatalf("PkgName: %v Expected: %v, got: %v\n", uri, expected.PkgName, computed.PkgName)
		}
		if computed.Version != expected.Version {
			t.Fatalf("Version: %v Expected: %v, got: %v\n", uri, expected.Version, computed.Version)
		}
	}
}

func TestExamineSecond(t *testing.T) {
	t.Skip()
}

func TestStripURI(t *testing.T) {
	stripTests := make(map[string]string)

	stripTests["https://pypi.python.org/packages/ec/51/df65c73f957375bbd814bb0353213d0422d07fe538b92a2a23fa683e499b/meson-0.33.0.tar.gz#md5=3252395e6df14e6f85270abb2542e49b"] = "https://pypi.python.org/packages/ec/51/df65c73f957375bbd814bb0353213d0422d07fe538b92a2a23fa683e499b/meson-0.33.0.tar.gz"
	stripTests["https://gitlab.com/manaplus/manaplus/repository/archive.zip?ref=v1.6.7.30"] = "https://gitlab.com/manaplus/manaplus/repository/archive.zip?ref=v1.6.7.30"
	stripTests["https://pypi.python.org/packages/fc/f1/7530ac8594453fc850e53580256f3152a8d8f2bb351bc3d0df8d7b53dbde/ruamel.yaml-0.11.11.tar.gz"] = "https://pypi.python.org/packages/fc/f1/7530ac8594453fc850e53580256f3152a8d8f2bb351bc3d0df8d7b53dbde/ruamel.yaml-0.11.11.tar.gz"

	for uri, expected := range stripTests {
		ret, err := StripURI(uri)
		if err != nil {
			t.Fatalf("Fatal: cannot strip %s: %v\n", uri, err)
		}
		if ret != expected {
			t.Fatalf("Fatal: %s does not match %s\n", ret, expected)
		}
	}
}
