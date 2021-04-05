package version

import (
	"log"
	"testing"
)

func TestReactNative(t *testing.T) {
	version := GetVersion{
		RootPath: "../tests/react",
		Logger:   log.Default(),
	}

	info, found := version.isReactNative()
	if !found {
		t.Errorf("Path does not have a React Native / Expo app.")
		return
	}

	if info == nil {
		t.Errorf("Info is nil")
		return
	}

	if info.Technology != ReactNative {
		t.Errorf("Invalid techonology. Excepted '%v'. Got '%v'", ReactNative, info.Technology)
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

	if info.File != "../tests/react/app.json" {
		t.Errorf("Invalid version. Excepted '../tests/react/app.json'. Got %s", info.File)
		return
	}
}
