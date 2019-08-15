package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Commands
/*
	cat *pathname*
		print file contents to command line
	cd *pathname*
		current dir set to pathname
	echo *words*
		words is echo'd to stdout
	exit *status*
	search *pathname* -word-
		searches file's content at *pathname* for word and maybe returns line number or something
	ls *pathname*
		files and dirs at *pathname* are printed
	make *pathname* [words...]
	mkdir *pathname*
	pwd
	rm *pathname
*/

func main() {

	// args := os.Args[1:]

	switch os.Args[1] {

	case "cat":
		path := os.Args[2]
		err := FileRead(path)
		fmt.Println(err)

	case "make":
		path := os.Args[2]
		content := os.Args[3:]
		joined := strings.Join(content, " ")
		err := FileWrite(path, joined)
		fmt.Println(err)
	}
}

func FileRead(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		return err
	}
	fmt.Println("File: %s is %d bytes long", path, fi.Size())

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return nil
}

func FileWrite(path string, content string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	n1, err := f.WriteString(content)
	if err != nil {
		return err
	}
	fmt.Printf("wrote %d bytes\n", n1)

	f.Sync()

	return nil
}
