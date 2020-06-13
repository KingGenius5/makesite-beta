package main

import (
	//"fmt"
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

type fluff struct {
	Content string
}

func main() {
	c := flag.String("file", "first-post.txt", "txt file to be read.")
	flag.Parse()

	post := readFile(*c)

	newName := strings.Split(*c, ".")[0] + ".html"

	writeFile(newName, post)
}

func readFile(fileName string) string {
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func writeFile(fileName string, text string) {
	paths := []string{
		"template.tmpl",
	}

	t := template.Must(template.New("template.tmpl").ParseFiles(paths...))
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	err = t.Execute(file, text)
	if err != nil {
		panic(err)
	}
}
