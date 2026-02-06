package downloader

type flagConfigItem struct {
	Name         string
	DefaultValue string
	Shorthand    string
	Usage        string
}

type flags struct {
	File   flagConfigItem
	Output flagConfigItem
}

var FlagsConfig = flags{
	File: flagConfigItem{
		Name:         "file",
		DefaultValue: "images.txt",
		Shorthand:    "f",
		Usage:        "Path to the text file containing image URLs",
	},
	Output: flagConfigItem{
		Name:         "output",
		DefaultValue: "images",
		Shorthand:    "o",
		Usage:        "Directory to save downloaded images",
	},
}

type flagValues struct {
	filePath  string
	outputDir string
}
