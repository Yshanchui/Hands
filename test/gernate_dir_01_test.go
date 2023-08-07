package test

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const jsonFileName = "dir.json"

var rootDir string
var separator string
var jsonData map[string]interface{}

func loadJson() {
	separator = string(filepath.Separator)
	workDir, _ := os.Getwd()
	rootDir = workDir[:strings.LastIndex(workDir, separator)]

	jsonBytes, _ := os.ReadFile(workDir + separator + jsonFileName)
	err := json.Unmarshal(jsonBytes, &jsonData)
	if err != nil {
		panic("Load Json Data Error: " + err.Error())
	}
}

func parseMap(mapData map[string]interface{}, parentDir string) {
	for k, v := range mapData {
		switch v.(type) {
		case string:
			{
				path, _ := v.(string)
				if path == "" {
					continue
				}
				if parentDir != "" {
					path = parentDir + separator + path
					if k == "text" {
						parentDir = path
					}
				} else {
					parentDir = path
				}
				createDir(path)
			}
		case []interface{}:
			{
				parseArray(v.([]interface{}), parentDir)
			}

		}
	}
}

func parseArray(jsonData []interface{}, parentDir string) {
	for _, v := range jsonData {
		mapV, _ := v.(map[string]interface{})
		parseMap(mapV, parentDir)
	}
}

func createDir(path string) {
	if path == "" {
		return
	}
	fmt.Println(path)
	err := os.MkdirAll(rootDir+separator+path, fs.ModePerm)
	if err != nil {
		panic("Create Dir Error: " + err.Error())
	}
}

func TestGernateDir01(t *testing.T) {
	loadJson()
	parseMap(jsonData, "")
}
