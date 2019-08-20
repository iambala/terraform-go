package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func cleanupFile() {
	input, err := ioutil.ReadFile("main.tf")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := bytes.Replace(input, []byte("$$"), []byte("$"), -1)

	if err = ioutil.WriteFile("main.tf", output, 0666); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
