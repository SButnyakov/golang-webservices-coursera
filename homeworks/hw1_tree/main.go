package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
)

func dirTreeRec(buffer *bytes.Buffer, dirName string, withFiles bool, prefix string) error {
	dir, err := os.Open(dirName)
	if err != nil {
		return err
	}

	filesArr, err := dir.Readdir(0)
	if err != nil {
		return err
	}

	files := make([]os.FileInfo, 0)

	for _, file := range filesArr {
		if file.IsDir() || withFiles {
			files = append(files, file)
		}
	}

	for i, file := range files {
		if file.Name() == ".idea" || file.Name() == "hw1.md" {
			continue
		}
		var str string
		if i == len(files)-1 && (file.IsDir() || (!file.IsDir() && withFiles)) {
			str += prefix + "└───" + file.Name()
		} else if file.IsDir() || (!file.IsDir() && withFiles) {
			str += prefix + "├───" + file.Name()
		}

		if !file.IsDir() {
			if file.Size() == 0 {
				str += " (empty)"
			} else {
				str += " (" + strconv.FormatInt(file.Size(), 10) + "b)"
			}
		}

		str += "\n"

		buffer.WriteString(str)

		if file.IsDir() {
			if i == len(files)-1 {
				err = dirTreeRec(buffer, dirName+string(os.PathSeparator)+file.Name(), withFiles, prefix+"\t")

			} else {
				err = dirTreeRec(buffer, dirName+string(os.PathSeparator)+file.Name(), withFiles, prefix+"│\t")
			}
		}
	}

	return err
}

func dirTree(buffer *bytes.Buffer, dirName string, withFiles bool) error {
	return dirTreeRec(buffer, dirName, withFiles, "")
}

func main() {
	withFilesPtr := flag.Bool("f", false, "flag for including files")
	flag.Parse()

	out := new(bytes.Buffer)
	err := dirTree(out, "testdata", *withFilesPtr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(out)
}
