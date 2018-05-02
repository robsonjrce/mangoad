package config

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func split(s string, c string) (string, string) {
	i := strings.Index(s, c)
	if i < 0 {
		return "", ""
	}
	return s[:i], s[i+1:]
}

func GetConfigContent() (err error) {
	var file *os.File

	if file, err = os.Open("./mangoad.conf"); err != nil {
		panic(err.Error())
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var line string
	for {
		line, err = reader.ReadString('\n')

		urlType, url := split(line, " ")

		println(urlType)
		println(url)

		if err != nil {
			break
		}
	}

	if err != io.EOF {
		fmt.Printf(" > Failed!: %v\n", err)
	}

	return
}
