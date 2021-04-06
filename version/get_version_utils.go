package version

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

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

func (v *GetVersion) getFile(filePath string) (string, *os.File, error) {
	var err error

	if _, err = os.Stat(filePath); err != nil {
		return filePath, nil, err
	}

	var file *os.File
	if file, err = os.Open(filePath); err != nil {
		return filePath, nil, err
	}

	return 	filePath, file, err
}

func (v *GetVersion) readFileContent(fileName string) (string, []byte, error) {

	var file *os.File
	var filePath string
	var err error

	fullFilePath := v.fullPathForFile(fileName)

	filePath, file, err = v.getFile(fullFilePath)
	if err != nil {
		return filePath, nil, err
	}

	defer file.Close()

	var content []byte
	if content, err = ioutil.ReadAll(file); err != nil {
		return filePath, nil, err
	}
	return filePath, content, err
}

func searchForFiles(root, pattern string) ([]string, error) {
	var matches []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if matched, err := filepath.Match(pattern, filepath.Base(path)); err != nil {
			return err
		} else if matched {
			matches = append(matches, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return matches, nil
}

func (v *GetVersion) fullPathForFile(filename string) string {
	return path.Join(v.RootPath, filename)
}
