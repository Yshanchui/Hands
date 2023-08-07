package test

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var rootDir02 string
var separator02 string

type Node struct {
	Text     string `json:"text"`
	Children []Node `json:"children"`
}

var rootNode Node

func loadJson02() {
	separator02 = string(filepath.Separator)
	workDir, _ := os.Getwd()
	rootDir02 = workDir[:strings.LastIndex(workDir, separator02)]

	jsonBytes, _ := os.ReadFile(workDir + separator02 + jsonFileName)
	err := json.Unmarshal(jsonBytes, &rootNode)
	if err != nil {
		panic("Load Json Data Error: " + err.Error())
	}
}

func parseNode(node Node, parentDir string) {
	if node.Text != "" {
		creatreDir02(node, parentDir)
	}
	if parentDir != "" {
		parentDir = parentDir + separator02
	}
	if node.Text != "" {
		parentDir = parentDir + node.Text
	}

	for _, v := range node.Children {
		parseNode(v, parentDir)
	}
}

func creatreDir02(node Node, parentDir string) {
	dirPath := rootDir02 + separator02
	if parentDir != "" {
		dirPath = dirPath + parentDir + separator02
	}
	dirPath = dirPath + node.Text
	err := os.MkdirAll(dirPath, fs.ModePerm)
	if err != nil {
		panic("Create Dir Error: " + err.Error())

	}
}
func TestGernateDir02(t *testing.T) {
	loadJson02()
	parseNode(rootNode, "")
}
