package main

import (
	"github.com/russross/blackfriday"
	"fmt"
  "os"
	"io/ioutil"
)

const usage = `
usage: ./mark2h [-h] filePath

flags
 h : show this message.

commands
 filePath: convert file path
`

func showUsage(exitCode int) {
	fmt.Println(usage)
	os.Exit(exitCode)
}

func main(){
	if len(os.Args) == 1 {
		showUsage(1)
	}

	if os.Args[1] == "-h" || os.Args[1] == "--help" {
		showUsage(0)
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	html := blackfriday.MarkdownBasic(data)
	fmt.Println(string(html))
}
