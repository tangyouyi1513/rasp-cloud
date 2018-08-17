package models

import (
	"os"
	"rasp-cloud/tools"
	"fmt"
	"crypto/md5"
)

type Plugin struct {
	Version string `json:"version"`
	Md5     string `json:"md5,omitempty"`
	Content string `json:"plugin,omitempty"`
}

var (
	PluginPrefix = "official-"
)

func init() {
	if isExists, _ := tools.PathExists("plugin"); !isExists {
		err := os.Mkdir("plugin", os.ModePerm)
		if err != nil {
			tools.Panic(err.Error())
		}
	}
}

func GetLatestPluginFromDir() (plugin *Plugin, err error) {
	jsFiles, err := tools.ListFiles("plugin", "js", PluginPrefix)
	if err != nil {
		return
	}
	if len(jsFiles) > 0 {
		newVersion := jsFiles[0][len(PluginPrefix) : len(jsFiles[0])-3]
		fileContent, readErr := tools.ReadFromFile("plugin/" + jsFiles[0])
		if readErr != nil {
			err = readErr
			return
		}
		plugin = NewPlugin(newVersion, fileContent)
	}
	return
}

func NewPlugin(version string, content []byte) *Plugin {
	newMd5 := fmt.Sprintf("%x", md5.Sum(content))
	return &Plugin{Version: version, Md5: newMd5, Content: string(content)}
}
