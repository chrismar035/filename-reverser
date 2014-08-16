package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"strings"
)

func reverse(normal string) string {
	reversed := ""
	chars := strings.Split(normal, "")
	for i := 0; i < len(chars); i++ {
		reversed = chars[i] + reversed
	}
	return reversed
}

func main() {
	var targetPath = "."
	if len(os.Args) > 1 {
		targetPath = os.Args[1]
	}

	fmt.Println("Procesing " + targetPath)

	files, err := ioutil.ReadDir(targetPath)
	if err != nil {
		fmt.Println("Unable to read from " + targetPath)
		fmt.Println(err)
		return
	}

	fmt.Println("")
	for _, file := range files {
		splits := strings.Split(file.Name(), ".")
		extension := splits[len(splits) - 1]
		reversedName := ""
		if len(splits) > 2 {
			for i := 0; i < len(splits) - 1; i ++ {
				reversedName = reversedName + splits[i]
			}
		} else {
			reversedName = splits[0]
		}

		fmt.Println(reversedName)
		fmt.Println(extension)

		correctedName := reverse(reversedName)
		originalFile := reversedName +  "." + extension
		correctedFile := correctedName +  "." + extension
		if targetPath != "." {
			originalFile = targetPath + originalFile
			correctedFile = targetPath + correctedFile
		}

		err := os.Rename(originalFile, correctedFile)

		if err != nil {
			fmt.Println("Coun't rename " + originalFile)
			fmt.Println(err)
		} else {
			fmt.Println(originalFile + " => " + correctedFile)
		}
	}
}
