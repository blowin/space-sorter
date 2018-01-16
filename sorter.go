// main project main.go
package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const (
	sep = string(os.PathSeparator)
)

func Sort(path string) *sync.WaitGroup {
	filePath := path + sep
	workDir, err := ioutil.ReadDir(filePath)
	if err != nil {
		return nil
	}
	var waiter sync.WaitGroup
	for _, dir := range workDir {
		waiter.Add(1)
		go dirSorter(filePath, dir, &waiter)
	}

	return &waiter
}

func nameSorter(filePath string, dir os.FileInfo, waiter *sync.WaitGroup,
	flag FExt) {
	defer waiter.Done()
	if dir.IsDir() {
		return
	}

}

func dirSorter(filePath string, dir os.FileInfo, waiter *sync.WaitGroup) {
	defer waiter.Done()
	if dir.IsDir() {
		return
	}

	ext, err := getExtencion(filePath + dir.Name())
	fileName := dir.Name()
	absFP := filePath + fileName
	if err {
		return
	}
	createDirPath := filePath + ext
	if err := os.MkdirAll(createDirPath, os.ModePerm); err != nil {
		return
	}
	newFilePath := filePath + ext + sep + fileName
	os.Rename(absFP, newFilePath)
}

func getExtencion(absFilePath string) (string, bool) {
	extenc := filepath.Ext(absFilePath)
	extLen := len(extenc)
	if extLen > 0 {
		return strings.ToLower(extenc[1:extLen]), false
	}
	return "", true
}
