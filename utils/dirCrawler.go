package utils

import (
	"log"
	"os"
)

func DirCrawler(path string, onFile func(string, os.DirEntry), onDir func(string, os.DirEntry)) {
	dirContent, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range dirContent {
		if file.IsDir() {
			onDir(path, file)
			DirCrawler(path+"/"+file.Name(), onFile, onDir)
		} else {
			onFile(path, file)
		}
	}
}
