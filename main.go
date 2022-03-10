package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
)

func main() {

	var (
		isDecode bool
		file     string
	)
	flag.BoolVar(&isDecode, "d", false, "decode data")
	flag.StringVar(&file, "f", "", "filepath, with no [f], read standard input")

	flag.Parse()

	var input string
	var output string

	if file == "" {
		if _, err := fmt.Scan(&input); err != nil {
			fmt.Println(err)
		}
	} else {
		tmpByte, _ := os.ReadFile(file)
		input = string(tmpByte)
	}

	if isDecode {
		tmpByte, _ := base64.StdEncoding.DecodeString(input)
		output = string(tmpByte)
	} else {
		output = base64.StdEncoding.EncodeToString([]byte(input))
	}

	fmt.Println(output)
}
