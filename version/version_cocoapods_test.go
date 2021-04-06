package version

import (
	"log"
	"testing"
)

func TestCocoapods(t *testing.T) {
	version := GetVersion{
		RootPath: "../tests/cocoapods",
		Logger:   log.Default(),
	}

	info, err := version.isCocoapods()
	if err != nil {
		t.Errorf("Path does not have a Cocoapods project.")
		return
	}

	if info == nil {
		t.Errorf("Info is nil")
		return
	}

	if info.Technology != Cocoapods {
		t.Errorf("Invalid techonology. Excepted '%v'. Got '%v'", Cocoapods, info.Technology)
		return
	}

	if info.Version != "1.2.3" {
		t.Errorf("Invalid version. Excepted '1.2.3'. Got %s", info.Version)
		return
	}

	if info.AppName != "Reachability" {
		t.Errorf("Invalid version. Excepted 'Reachability'. Got %s", info.AppName)
		return
	}

	if info.File != "../tests/cocoapods/Reachability.podspec" {
		t.Errorf("Invalid version. Excepted '../tests/cocoapods/Reachability.json'. Got %s", info.File)
		return
	}
}
