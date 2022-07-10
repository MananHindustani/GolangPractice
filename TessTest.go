package main

import (
	"github.com/cavo2/commons"
	"log"
	"os"
	"path/filepath"
	"fmt"
	"io/ioutil"
)

const ppmfilespath = "/go/test/Normal20062453_H1415_IL_160361650/rotated/"

func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func OSReadDir(root string) ([]string, error) {
	var files []string
	f, err := os.Open(root)
	if err != nil {
		return files, err
	}
	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}

func IOReadDir(root string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}

func WalkFn() (ppms []string) {
	var (
		files []string
		err error
	)


	// filepath.Walk
	files, err = FilePathWalkDir(ppmfilespath)
	if err != nil {
		panic(err)
	}
	// ioutil.ReadDir
	files, err = IOReadDir(ppmfilespath)
	if err != nil {
		panic(err)
	}
	//os.File.Readdir
	files, err = OSReadDir(ppmfilespath)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}
	ppms = files
	return files
}

func main() {
	textFilePath := "/go/test/Normal20062453_H1415_IL_160361650/"
	ppms := WalkFn()
	f, err := os.Create(textFilePath + "/text.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	for _, ppm := range ppms {
		ppm = ppmfilespath + ppm
		jpeg,_:=commons.ImageToJPEG(ppm)
		fmt.Println(jpeg)
		txt, _, hocrFile, err := commons.Tesseract(ppm, 300)
		fmt.Println(hocrFile)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.WriteString(txt + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}

}