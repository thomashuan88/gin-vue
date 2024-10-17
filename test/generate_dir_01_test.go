package test

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

var stRootDir string
var stSeparator string
var iJsonData map[string]any

const stJsonFileName = "dir.json"

func loadJson(t *testing.T) {
	stSeparator = string(os.PathSeparator)
	stWorkDir, _ := os.Getwd()

	stRootDir = stWorkDir[:strings.LastIndex(stWorkDir, stSeparator)]

	gnJsonBytes, _ := os.ReadFile(stWorkDir + stSeparator + stJsonFileName)

	err := json.Unmarshal(gnJsonBytes, &iJsonData)

	if err != nil {
		t.Error("Load Json Data Error: " + err.Error())
	}

}

func parseMap(t *testing.T, mapData map[string]any, stParentDir string) {
	for k, v := range mapData {
		switch v.(type) {
		case string:
			// create dir
			path, _ := v.(string)
			if path == "" {
				continue
			}

			if stParentDir != "" {
				path = stParentDir + stSeparator + path
				if k == "text" {
					stParentDir = path
				}
			} else {
				stParentDir = path
			}
			createDir(path)
		case []any:
			parseArray(t, v.([]any), stParentDir)

		}
	}

}

func parseArray(t *testing.T, arrayData []any, stParentDir string) {
	for _, v := range arrayData {
		mapV, _ := v.(map[string]any)
		parseMap(t, mapV, stParentDir)
	}
}

func createDir(path string) {
	if path == "" {
		return
	}
	err := os.MkdirAll(stRootDir+stSeparator+path, os.ModePerm)

	if err != nil {
		panic(err)
	}
}

func TestGenerateDir01(t *testing.T) {
	loadJson(t)
	parseMap(t, iJsonData, "")
}
