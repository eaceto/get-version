package version

func (v *GetVersion) isReactNative() (*Info, error) {

	const appJSONFileName = "app.json"
	filePath, packageJson, err := v.readJSONFile(appJSONFileName)

	if err != nil {
		return nil, err
	}

	expo := packageJson["expo"].(map[string]interface{})
	if expo == nil {
		return nil, err
	}

	return &Info{
		Version:    expo["version"].(string),
		AppName:    expo["name"].(string),
		Technology: ReactNative,
		File:       filePath,
	}, nil
}
