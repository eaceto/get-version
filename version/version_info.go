package version

type Technology string
const (
	JavaScript  Technology = "JavaScript / TypeScript"
	ReactNative            = "ReactNative / ReactNative Native"
	Maven	=  "Maven"
	Gradle = "Gradle"
)

type Info struct {
	AppName string
	Version string
	Technology Technology
	File string
}
