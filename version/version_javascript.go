package version

func (v *GetVersion) isJavascript() (*Info, error) {

	const packageJsonFileName = "package.json"
	filePath, packageJson, err := v.readJSONFile(packageJsonFileName)

	if err != nil {
		return nil, err
	}

	return &Info{
		Version:    packageJson["version"].(string),
		AppName:    packageJson["name"].(string),
		Technology: JavaScript,
		File:       filePath,
	}, nil
}
