package main

import (
	//"fmt"

	"html/template"
	"io/ioutil"
	"os"
)

type fluff struct {
	Content string
}

func main() {
	// Files are provided as a slice of strings.
	paths := []string{
		"template.tmpl",
	}

	post := readFile()

	t := template.Must(template.New("template.tmpl").ParseFiles(paths...))
	file, err := os.Create("first-post.html")
	if err != nil {
		panic(err)
	}

	err = t.Execute(file, post)
	if err != nil {
		panic(err)
	}
}

func readFile() string {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		panic(err)
	}

	return string(fileContents)
}

func writeFile(text string) {
	bytesToWrite := []byte(text)
	err := ioutil.WriteFile("first-post.html", bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
	return
}
