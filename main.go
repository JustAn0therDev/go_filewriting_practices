package main

import (
	"fmt"
	"log"
	"os"
	/*
		"bufio"
		"os"
	*/)

func check(e error) {
	if e != nil {
		// the built-in "panic" function stops all of the 
		// call stack execution, working like a "throw exception".
		// When it gets to the beginning of the call stack, the program is terminated
		panic(e)
	}
}

func main() {
	writeContentToFile()
	// appendContentToFile()
	readFileContent()
}

func writeContentToFile() {
	var err error
	newFileContent := []byte("Changed the content of the file!")
	file, err := os.OpenFile("./hello.txt", os.O_RDWR, os.ModeAppend)

	check(err)

	file.Seek(1, 0) // this function call should keep the "C" at the beginning of the file

	bytesWritten, err := file.Write(newFileContent)

	fmt.Printf("Number of written bytes: %v \n", bytesWritten)
	
	check(err)

	file.Close()
}

func appendContentToFile() {
	var err error
	newFileContent := []byte("Added content to file!")
	file, err := os.OpenFile("./hello.txt", os.O_APPEND, os.ModeAppend)

	check(err)

	bytesWritten, err := file.Write(newFileContent)

	fmt.Printf("Number of written bytes: %v \n", bytesWritten)
	
	check(err)

	file.Close()
}

func readFileContent() {
	bytesFromFile, err := os.ReadFile("hello.txt")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bytesFromFile))
}