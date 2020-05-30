package drop

import (
	"io/ioutil"
	"os"
)

// Dropit drops down directories
func Dropit(path string) map[string]bool {
	isItADir, _ := os.Stat(path)
	subDirList := make(map[string]bool)
	if isItADir.IsDir() {
		dir, _ := ioutil.ReadDir(path)
		for _, file := range dir {
			if !file.IsDir() {
				subDirList[file.Name()] = false
			} else {
				subDirList[file.Name()] = true
			}
		}
	}
	return subDirList
}
