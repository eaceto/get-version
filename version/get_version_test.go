package version

import (
	"log"
	"testing"
)

func TestValidPath(t *testing.T) {
	version := GetVersion{
		RootPath: "../tests/maven",
		Logger: log.Default(),
	}

	info, err :=version.Analyze()
	if err != nil {
		t.Errorf("fail to analyze path. %v", err)
	}

	if info == nil {
		t.Errorf("info is nil")
	}
}
