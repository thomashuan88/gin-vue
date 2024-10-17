package test

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

type Node struct {
	Text     string `json:"text"`
	Children []Node `json:"children"`
}

var stRootDir2 string
var stSeparator2 string
var stJsonFileName2 = "dir.json"
var iRootNode Node

func loadJson2(t *testing.T) {
	stSeparator2 = string(os.PathSeparator)
	stworkDir, _ := os.Getwd()
	stRootDir2 = stworkDir[:strings.LastIndex(stworkDir, stSeparator2)]

	gnJsonBytes, _ := os.ReadFile(stworkDir + stSeparator2 + stJsonFileName2)
	err := json.Unmarshal(gnJsonBytes, &iRootNode)
	if err != nil {
		t.Error("Load Json Data Error: " + err.Error())
	}
}

func parseNode(t *testing.T, iNode *Node, stParentDir string) {
	if iNode.Text != "" {
		createDir2(iNode, stParentDir)
	}
	if stParentDir != "" {
		stParentDir = stParentDir + stSeparator2 + iNode.Text
	} else {
		stParentDir = iNode.Text
	}
	if len(iNode.Children) > 0 {
		for _, v := range iNode.Children {
			parseNode(t, &v, stParentDir)
		}
	}

}

func createDir2(iNode *Node, stParentDir string) {
	stDirPath := stRootDir2
	if stParentDir != "" {
		stDirPath = stDirPath + stSeparator2 + stParentDir
	}

	stDirPath = stDirPath + stSeparator2 + iNode.Text

	err := os.MkdirAll(stDirPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

func TestGenerateDir02(t *testing.T) {
	loadJson2(t)
	parseNode(t, &iRootNode, "")
}
