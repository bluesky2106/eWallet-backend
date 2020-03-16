package config

import (
	"fmt"
	"strings"
)

func getFileNameAndType(file string) (fileName string, fileType string) {
	strs := strings.Split(file, ".")

	if len(strs) != 2 {
		fmt.Println("Config file name must follow format of xxx.yyy")
		return "", ""
	}

	fileName = strs[0]
	fileType = strs[1]

	return fileName, fileType
}
