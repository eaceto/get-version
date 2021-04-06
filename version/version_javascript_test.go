package version

import (
	"log"
	"testing"
)

func TestJavaScript(t *testing.T) {
	version := GetVersion{
		RootPath: "../tests/nodejs",
		Logger:   log.Default(),
	}

	info, err := version.isJavascript()
	if err != nil {
		t.Errorf("Path does not have a JavaScript / TypeScript app.")
		return
	}

	if info == nil {
		t.Errorf("Info is nil")
		return
	}

	if info.Technology != JavaScript {
		t.Errorf("Invalid techonology. Excepted '%v'. Got '%v'", JavaScript, info.Technology)
		return
	}

	if info.Version != "1.2.3" {
		t.Errorf("Invalid version. Excepted '1.2.3'. Got %s", info.Version)
		return
	}

	if info.AppName != "example-nodejs-app" {
		t.Errorf("Invalid version. Excepted 'example-nodejs-app'. Got %s", info.AppName)
		return
	}

	if info.File != "../tests/nodejs/package.json" {
		t.Errorf("Invalid version. Excepted '../tests/nodejs/package.json'. Got %s", info.File)
		return
	}
}
