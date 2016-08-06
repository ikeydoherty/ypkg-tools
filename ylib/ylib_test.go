package ylib

import (
	"testing"
)

func TestExamineURI(t *testing.T) {
	// url := args[1]
	if info := ExamineURI("https://github.com/solus-project/linux-steam-integration/releases/download/v0.2/linux-steam-integration-0.2.tar.xz"); info != nil {
		if info.PkgName != "linux-steam-integration" {
			t.FailNow()
		}
		if info.Version != "0.2" {
			t.FailNow()
		}
	} else {
		t.FailNow()
	}

	// github v URL
	if info := ExamineURI("https://github.com/solus-project/linux-steam-integration/archive/v0.2.tar.gz"); info != nil {
		if info.PkgName != "linux-steam-integration" {
			t.FailNow()
		}
		if info.Version != "0.2" {
			t.FailNow()
		}
	} else {
		t.FailNow()
	}

	// pypi url with odd name package
	if info := ExamineURI("https://pypi.python.org/packages/fc/f1/7530ac8594453fc850e53580256f3152a8d8f2bb351bc3d0df8d7b53dbde/ruamel.yaml-0.11.11.tar.gz"); info != nil {
		if info.PkgName != "ruamel.yaml" {
			t.FailNow()
		}
		if info.Version != "0.11.11" {
			t.FailNow()
		}
	} else {
		t.FailNow()
	}

	// Very odd sourceforge URL
	if info := ExamineURI("http://internode.dl.sourceforge.net/project/yodl/yodl/3.05.01/yodl_3.05.01.orig.tar.gz"); info != nil {
		if info.PkgName != "yodl" {
			t.FailNow()
		}
		if info.Version != "3.05.01.orig" {
			t.FailNow()
		}
	} else {
		t.FailNow()
	}
}

func TestExamineSecond(t *testing.T) {
	t.Skip()
}
