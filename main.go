package main

import (
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
	//reader := bufio.NewReader(os.Stdin)
	fmt.Println("Select a framework")
	for i := range frameworks {
		fmt.Printf("%d) %s\n", i, frameworks[i].Name)
	}
	fmt.Print("\n> ")
	var selectedFramework int
	_, err := fmt.Scanf("%d", &selectedFramework)
	if err != nil {
		panic(err)
	}
	os.Mkdir("."+string(filepath.Separator)+frameworks[selectedFramework].Name, 0755)

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

// read strings
//nameArr, _, _ := reader.ReadLine()
//name := string(nameArr)
