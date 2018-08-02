package tools

import (
	"os"
	"strings"
	"io/ioutil"
	"sort"
)

func ReadFromFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return data, err
	}

	return data, nil
}

func ListFiles(dirPath string, suffix string, prefix string) (files []string, err error) {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		if os.IsPermission(err) {
			err = os.Chmod("plugin", os.ModePerm)
			if err == nil {
				dir, err = ioutil.ReadDir(dirPath)
			}
		}
		if err != nil {
			return
		}
	}

	for _, file := range dir {
		if !file.IsDir() {
			if strings.HasSuffix(file.Name(), suffix) && strings.HasPrefix(file.Name(), prefix) {
				files = append(files, file.Name())
			}
		}
	}

	sort.Sort(sort.Reverse(sort.StringSlice(files)))
	return files, nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
