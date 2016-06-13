package main

import (
	"github.com/russross/blackfriday"
	"fmt"
  "os"
	"io/ioutil"
	"html/template"
)

const usage = `
usage: ./mark2h [-sh] filePath

flags
 s : add style
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
	var style bool

	if os.Args[1] == "-s" || os.Args[1] == "--style" {
		// TODO: select style
		style = true
		filePath = os.Args[2]
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	html := blackfriday.MarkdownCommon(data)

	// Print simple HTML
	if !style {
		fmt.Println(string(html))
		os.Exit(0)
	}

	tpl, err := ioutil.ReadFile("markdown.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	t := template.Must(template.New("markdown").Parse(string(tpl)))
	if err = t.Execute(os.Stdout, template.HTML(string(html))); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
