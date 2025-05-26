package main

import (
	"fmt"
	"os"
	"github.com/BurntSushi/toml"
)

type Config struct {
	// struct to store the configuration
	Version string
	LocalStorage bool
	StorageAddress string
}



func createNoteString(words []string) string {
	// parse command line arguments into a single string
	var s, sep string

	for _, val := range words {
		s += sep + val
		sep = " "
	}
	return s
}

func saveNoteString(noteString []string, conf Config) (string, error) {
	//tbd

	return "Success!", nil
}

func countNotes() {
	//tbd
}

func listNotes() {
	//tbd
}

func removeNote() {}


func main() {
	// first load up the config.toml file
	var conf Config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		fmt.Println(err)
		
	}

	// pull in all the args first
	cmd := os.Args[1]
	fmt.Printf("---Taskr Version %s---\n", conf.Version)
	
	fmt.Printf("You selected '%s' command\n", cmd)

	var result string
	var err error


	if cmd == "add" {
		
		result, err = saveNoteString(os.Args[2:], conf)

	}


	// return the results to the terminal
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	

}


