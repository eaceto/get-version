package version

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func (v *GetVersion) isCocoapods() (*Info, error) {

	files, err := searchForFiles(v.RootPath, "*.podspec")

	if err != nil {
		return nil, err
	}

	if len(files) != 1 {
		return nil, fmt.Errorf("invalid number of podspec files found. Expected '1', got: %d", len(files))
	}

	var pod *podspec
	var filePath string

	filePath, pod, err = v.readPodspecFile(files[0])

	return &Info{
		Version:    pod.version,
		AppName:    pod.name,
		Technology: Cocoapods,
		File:       filePath,
	}, nil
}

func (v *GetVersion) readPodspecFile(fileName string) (string, *podspec, error) {

	filePath, file, err := v.getFile(fileName)
	if err != nil {
		return filePath, nil, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	name := ""
	version := ""

	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}

		trimmedLine := strings.TrimSpace(line)

		if strings.HasPrefix(trimmedLine, "spec.name") ||
			strings.HasPrefix(trimmedLine, "spec.version") {
			val := strings.ReplaceAll(trimmedLine, "spec.name", "")
			val = strings.ReplaceAll(val, "spec.version", "")
			val = strings.ReplaceAll(val, "'", "")
			val = strings.ReplaceAll(val, "\"", "")
			val = strings.ReplaceAll(val, "=", "")
			val = strings.TrimSpace(val)

			if strings.HasPrefix(trimmedLine, "spec.name") {
				name = val
			}
			if strings.HasPrefix(trimmedLine, "spec.version") {
				version = val
			}
		}

		if len(name) > 0 && len(version) > 0 {
			break
		}
	}

	if len(version) > 0 && len(name) > 0 {
		return filePath, &podspec{
			name,
			version,
		}, nil
	}

	if len(version) == 0 {
		return filePath, nil, fmt.Errorf("missing version on podspec")
	}

	return filePath, nil, fmt.Errorf("missing name on podspec")
}
