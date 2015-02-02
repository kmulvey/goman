package main

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

type Framework struct {
	Name string
}

func main() {
	/*
		s := spin.New()
		for i := 0; i < 30; i++ {
			fmt.Printf("\r  \033[36mcomputing\033[m %s ", s.Next())
			time.Sleep(100 * time.Millisecond)
		}
	*/

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter project name: ")
	nameArr, _, _ := reader.ReadLine()
	name := string(nameArr)
	os.Mkdir("."+string(filepath.Separator)+name, 0755)

	p := Framework{name}

	// open output file
	fo, err := os.Create("test/server.go")
	if err != nil {
		panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	// process template
	tmpl, err := template.ParseFiles("templates/gin/server.go")
	_ = tmpl.ExecuteTemplate(fo, tmpl.Name(), p)
	if err != nil {
		panic(err)
	}

}
