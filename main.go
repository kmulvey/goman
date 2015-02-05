package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/tj/go-spin"
)

func main() {
	// load configs
	var frameworks []Framework
	getConfigs(&frameworks)

	// talk to the user
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
	// create the directory
	os.Mkdir("."+string(filepath.Separator)+frameworks[selectedFramework].Name, 0755)

	messages := make(chan string)
	go postScripts(messages)
	s := spin.New()
	var done string
	for done == "" {
		fmt.Printf("\r  \033[36mcomputing\033[m %s ", s.Next())
		done = <-messages
		fmt.Println("done: ", done)
		time.Sleep(100 * time.Millisecond)
	}
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
func postScripts(messages chan string) {
	cmd := exec.Command("go", "get")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	err = cmd.Wait()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
	time.Sleep(10000 * time.Millisecond)
	messages <- "done"
}

// read strings
//reader := bufio.NewReader(os.Stdin)
//nameArr, _, _ := reader.ReadLine()
//name := string(nameArr)

/*
spinner
s := spin.New()

		for i := 0; i < 30; i++ {
			fmt.Printf("\r  \033[36mcomputing\033[m %s ", s.Next())
			time.Sleep(100 * time.Millisecond)
		}
*/
