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
	fmt.Printf("$>")
	// args := os.Args[1:]
	scanner := bufio.NewScanner(os.Stdin)
	//scanner.Split()
	args := make([]string, 0)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		args = append(args, scanner.Text())
		fmt.Printf("$>")
		args = strings.Split(args[0], " ")

		if len(args) == 2 {
			switch args[0] {

			case "cat":
				path := args[1]
				err := FileRead(path)
				fmt.Println(err)
			}
		} else if len(args) > 2 {
			switch args[0] {

			case "make":
				path := args[1]
				content := args[2:]
				joined := strings.Join(content, " ")
				err := FileWrite(path, joined)
				fmt.Println(err)
			}
		}

	}
	fmt.Println("Here")
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
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
