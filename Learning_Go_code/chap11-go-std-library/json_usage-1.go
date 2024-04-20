package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func t1() {
	toFile := Person{
		"Fred",
		40,
	}
	tmpFile, err := ioutil.TempFile(os.TempDir(), "sample-")
	if err != nil {
		panic(err)
	}
	defer os.Remove(tmpFile.Name())
	err = json.NewEncoder(tmpFile).Encode(toFile)
	if err != nil {
		panic(err)
	}
	err = tmpFile.Close()
	if err != nil {
		panic(err)
	}

	// get file from tempfile
	tmpFile2, err := os.Open(tmpFile.Name())
	if err != nil {
		panic(err)
	}
	var fromFile Person
	err = json.NewDecoder(tmpFile2).Decode(&fromFile)
	if err != nil {
		panic(err)
	}

	err = tmpFile2.Close()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", fromFile)
}

func main() {
	t1()
}
