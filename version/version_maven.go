package version

import "encoding/xml"

func (v *GetVersion) isMaven() (*Info, error) {

	const pomFileName = "pom.xml"
	var filePath string
	var content []byte
	var err error

	filePath, content, err = v.readFileContent(pomFileName)

	if err != nil {
		return nil, err
	}

	var pom = mavenType{}

	if err = xml.Unmarshal(content, &pom); err != nil {
		return nil, err
	}


	return &Info{
		Version:    pom.Version,
		AppName:    pom.ArtifactId,
		Technology: Maven,
		File:       filePath,
	}, nil
}

type mavenType struct {
	XMLName xml.Name `xml:"project"`
	ArtifactId string `xml:"artifactId"`
	Version    string `xml:"version"`
}
