package version

import (
	"fmt"
	"log"
)

type GetVersion struct {
	RootPath string
	Logger   *log.Logger
}

type verifyLanguage func()(*Info,  error)


func (v *GetVersion)Analyze() (*Info, error) {

	infoFound := make([]*Info, 0)

	languages := []verifyLanguage {
		v.isReactNative,
		v.isJavascript,
		v.isMaven,
	}

	for _, language := range languages {
		if info, _ := language(); info != nil {
			infoFound = append(infoFound, info)
		}
	}

	infoCount := len(infoFound)

	if infoCount == 1 {
		return infoFound[0], nil
	}

	if infoCount > 1 {
		return nil, fmt.Errorf("multiple source codes declaration found at path: %s. Expected: 1, got: %d", v.RootPath, infoCount)
	}

	return nil, fmt.Errorf("no valid source code fount at path: %s", v.RootPath)
}
