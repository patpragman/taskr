package main

import (
	"encoding/csv"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	// struct to store the configuration
	Version        string
	LocalStorage   bool
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

func saveNoteString(noteString string, conf Config) (string, error) {

	// first check if the file exists
	fileExists := false
	if _, err := os.Stat(conf.StorageAddress); err == nil {
		fileExists = true
	}

	// next open the file for appending
	f, err := os.OpenFile(conf.StorageAddress, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}
	defer f.Close()

	// next create our writer and defer the flush until the function returns
	writer := csv.NewWriter(f)
	defer writer.Flush()

	// if the file is brand new you need to write headers
	if !fileExists {
		if err := writer.Write([]string{"Date", "Note"}); err != nil {
			return "", err
		}

	}

	// finally, we can write the actual row (man, you end up writing a lot of boilerplate in go)
	timestamp := time.Now().Format(time.RFC3339)
	if err := writer.Write([]string{timestamp, noteString}); err != nil {
		return "", err
	}

	return "Saved...", nil
}

func countNotes(conf Config) (string, error) {
	// open the csv file
	f, err := os.Open(conf.StorageAddress)
	if err != nil {
		return "0", err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return "0", err
	}

	return fmt.Sprintf("%d", len(records)), nil
}

func listNotes(conf Config) (string, error) {
	// open the .csv file and dump the indexed contents into the command line
	f, err := os.Open(conf.StorageAddress)
	if err != nil {
		return "", err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return "error reading file...", err
	}

	s := ""
	for i, record := range records {
		s += fmt.Sprintf("%d: (%s) %s\n", i+1, record[0], record[1])
	}

	return strings.TrimSuffix(s, "\n"), nil

}

func removeNote(i int, conf Config) (string, error) {
	// read the file then remove the ith row in the csv

	f, err := os.Open(conf.StorageAddress)
	if err != nil {
		return "", err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return "", err
	}

	// Check bounds
	if i < 1 || i > len(records) {
		return fmt.Sprintf("Invalid Index: %d", i), fmt.Errorf("invalid index: %d", i)
	}

	// now remove the row at that index...
	newRecords := append(records[:i-1], records[i:]...)

	// now we need to rewrite the file without the index
	f.Close()
	f, err = os.Create(conf.StorageAddress)
	if err != nil {
		return "Error updating file...", err
	}
	defer f.Close()

	writer := csv.NewWriter(f)
	defer writer.Flush()

	if err := writer.WriteAll(newRecords); err != nil {
		return "Error updating file...", err
	}

	return fmt.Sprintf("Removed the note %d", i), nil
}

func main() {

	// zeroth step, make sure you have the write path
	var err error
	var homeDir string

	// 1. Get the user's home directory
	homeDir, err = os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting user home directory: %v\n", err)
		os.Exit(1)
	}

	// 2. Construct the full, absolute path to the config file
	// Use filepath.Join to handle path separators correctly across OS
	configPath := filepath.Join(homeDir, ".local", "etc", "taskr", "config.toml")

	// first load up the config.toml file
	var conf Config
	if _, err = toml.DecodeFile(configPath, &conf); err != nil {
		fmt.Println(err)

	}

	// pull in all the args first
	var cmd string
	if len(os.Args) < 2 {
		cmd = "about"
	} else {
		cmd = os.Args[1]
	}

	//fmt.Printf("You selected '%s' command\n", cmd)

	var result string

	if cmd == "add" {
		result, err = saveNoteString(createNoteString(os.Args[2:]), conf)
	} else if cmd == "n" {
		result, err = countNotes(conf)
	} else if cmd == "list" {
		result, err = listNotes(conf)
	} else if cmd == "remove" || cmd == "rm" {

		i, err := strconv.Atoi(os.Args[2])
		if err == nil {
			result, err = removeNote(i, conf)
		}
	} else if cmd == "pop" {
		result, err = removeNote(1, conf)

	} else if cmd == "about" {
		// about and help stuff goes here
		result = `-----------------------------------------------
Taskr Version %s
By Pat Pragman, Pragman LLC
www.pragman.io
-----------------------------------------------
about -- takes you here
add <then text> -- adds that text to the note file
list -- lists all the current notes and their indices
remove <i> -- removes that particular index
n -- gives you a count of the current notes
-----------------------------------------------`
		result = fmt.Sprintf(result, conf.Version)
	}

	// return the results to the terminal
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
