package version

import (
	"log"
	"testing"
)

func TestMaven(t *testing.T) {
	version := GetVersion{
		RootPath: "../tests/maven",
		Logger:   log.Default(),
	}

	info, err := version.isMaven()
	if err != nil {
		t.Errorf("Path does not have a Maven script.")
		return
	}

	if info == nil {
		t.Errorf("Info is nil")
		return
	}

	if info.Technology != Maven {
		t.Errorf("Invalid techonology. Excepted '%v'. Got '%v'", Maven, info.Technology)
		return
	}

	if info.Version != "1.2.3" {
		t.Errorf("Invalid version. Excepted '1.2.3'. Got %s", info.Version)
		return
	}

	if info.AppName != "example-app" {
		t.Errorf("Invalid version. Excepted 'example-app'. Got %s", info.AppName)
		return
	}

	if info.File != "../tests/maven/pom.xml" {
		t.Errorf("Invalid version. Excepted '../tests/maven/pom.xml'. Got %s", info.File)
		return
	}
}
