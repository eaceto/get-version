package version

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
)

type GetVersion struct {
	RootPath string
	Logger   *log.Logger
}

type verifyLanguage func(info *Info,  bool)


func (v *GetVersion)Analyze() (*Info, error) {

	infoFound := make([]Info, 0)

	languages := []verifyLanguage {
		v.isReactNative,
		v.isJavascript,
		v.isMaven,
	}

	for i, language := range languages {

	}

	return nil, fmt.Errorf("no valid source code fount at path: %s", v.RootPath)
}

func (v *GetVersion) fullPathForFile(filename string) string {
	return path.Join(v.RootPath, filename)
}

func (v *GetVersion) readJSONFile(fileName string) (string, map[string]interface{}, error) {

	filePath, content, err := v.readFileContent(fileName)
	if err != nil {
		return filePath, nil, err
	}

	var packageJson map[string]interface{}
	if err := json.Unmarshal(content, &packageJson); err != nil {
		return filePath, nil, err
	}
	return filePath, packageJson, err
}

func (v *GetVersion) readFileContent(fileName string) (string, []byte, error) {
	var err error

	filePath := v.fullPathForFile(fileName)
	if _, err = os.Stat(filePath); err != nil {
		return filePath, nil, err
	}

	v.Logger.Printf("Found %s at %s", fileName, filePath)

	var jsonFile *os.File
	if jsonFile, err = os.Open(filePath); err != nil {
		return filePath, nil, err
	}

	var content []byte
	if content, err = ioutil.ReadAll(jsonFile); err != nil {
		return filePath, nil, err
	}
	return filePath, content, err
}
