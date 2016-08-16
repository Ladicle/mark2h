package main

import (
	"github.com/russross/blackfriday"
	"fmt"
  "os"
	"io/ioutil"
	"html/template"
)

const templatePath = "assets/markdown.html"
const usage = `
usage: ./mark2h [-sh] filePath

flags
 h : show this message

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

  filePath := os.Args[1]
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	html := blackfriday.MarkdownCommon(data)

	tpl, err := Asset(templatePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	t := template.Must(template.New(templatePath).Parse(string(tpl)))
	if err = t.Execute(os.Stdout, template.HTML(string(html))); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
