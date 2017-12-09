package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//some vars
	var filename string
	var collection map[string]int
	var index float32

	//get progs arguments
	args := os.Args

	if len(args) > 1 {
		filename = args[1]
	} else {
		fmt.Println("ERROR: No file given!")
		os.Exit(1)
	}

	fmt.Println("Start kappa-test\n")

	//init map
	collection = make(map[string]int)

	//read file
	content, err := ioutil.ReadFile(filename)
	con := string(content)

	//parse content and count letters
	if err != nil {
		log.Fatal(err)
	}
	for _, r := range con {
		//fmt.Println("read char "+strconv.Itoa(i))
		collection[strings.ToUpper(string(r))]++
	}
	//remove special chars, spaces and newline
	delete(collection, " ")
	delete(collection, "\t")
	delete(collection, "\n")
	delete(collection, "\r")
	delete(collection, ".")
	delete(collection, ";")
	delete(collection, ":")
	delete(collection, "'")
	delete(collection, "\"")
	delete(collection, "?")
	delete(collection, "$")
	delete(collection, "&")
	delete(collection, "%")
	delete(collection, "§")
	delete(collection, "-")
	delete(collection, "(")
	delete(collection, ")")
	delete(collection, "[")
	delete(collection, "]")

	//get numbers of letters (textlength without special chars ...)
	var length int
	for _, val := range collection {
		length += val
	}

	fmt.Println("Length of text (letters only): " + strconv.Itoa(length))

	fmt.Println("\nFrequency of letters")
	fmt.Println("| letter | num\t  | %\t   |")
	fmt.Println(strings.Repeat("-", 28))
	for key, val := range collection {
		f := float32(val) / float32(length)
		index += f * f
		fmt.Printf("| %s\t | %d\t  | %1.4f |\n", key, val, f)
	}

	fmt.Printf("\nIndex of coincidence: %1.4f\n\n", index)

	fmt.Println("done...")
}
