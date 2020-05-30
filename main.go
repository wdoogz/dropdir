package main

import (
	"droplog/drop"
	"fmt"
)

func main() {
	baseDir := "/var/log"
	for item, isDir := range drop.Dropit(baseDir) {
		fmt.Println(item, isDir)
		if !isDir {
			continue
		} else if isDir {
			fmt.Println(drop.Dropit(baseDir + "/" + item))
		}
	}
}
