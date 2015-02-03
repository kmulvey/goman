package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	/*
		s := spin.New()
		for i := 0; i < 30; i++ {
			fmt.Printf("\r  \033[36mcomputing\033[m %s ", s.Next())
			time.Sleep(100 * time.Millisecond)
		}
	*/
	// load configs
	var frameworks []Framework
	getConfigs(&frameworks)

	// talk to the user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter project name: ")
	nameArr, _, _ := reader.ReadLine()
	name := string(nameArr)
	os.Mkdir("."+string(filepath.Separator)+name, 0755)

	/*
		f := Framework{name, make([]string, 1)}

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
		_ = tmpl.ExecuteTemplate(fo, tmpl.Name(), f)
		if err != nil {
			panic(err)
		}
	*/
}

func getConfigs(frameworks *[]Framework) {
	file, e := ioutil.ReadFile("./frameworks.json")
	if e != nil {
		fmt.Printf("Unable to open config file: %v", e)
		os.Exit(1)
	}
	json.Unmarshal(file, &frameworks)
}
