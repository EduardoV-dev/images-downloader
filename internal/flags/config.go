// Package flags contains configuration and functions related cli flags
package flags

var FlagsConfig = flags{
	File: flagConfigItem[string]{
		Name:         "file",
		DefaultValue: "images.txt",
		Shorthand:    "f",
		Usage:        "Path to the text file containing image URLs",
	},
	Output: flagConfigItem[string]{
		Name:         "output",
		DefaultValue: "images",
		Shorthand:    "o",
		Usage:        "Directory to save downloaded images",
	},
	Timeout: flagConfigItem[uint]{
		Name:         "timeout",
		DefaultValue: 5,
		Shorthand:    "t",
		Usage:        "Timeout for fetching the images defined in seconds (time limit)",
	},
}

type flagConfigItem[T any] struct {
	Name         string
	DefaultValue T
	Shorthand    string
	Usage        string
}

type flags struct {
	File    flagConfigItem[string]
	Output  flagConfigItem[string]
	Timeout flagConfigItem[uint]
}

type flagValues struct {
	FilePath  string
	OutputDir string
	Timeout   uint
}
