package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Variables struct {
	Variables []Variable `json:"variables"`
}

type Variable struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func Config() {
	jsonFile, err := os.Open("./vars/env.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	var variables Variables
	json.Unmarshal(byteValue, &variables)
	for i := 0; i < len(variables.Variables); i++ {
		os.Setenv(variables.Variables[i].Name, variables.Variables[i].Value)
	}
	fmt.Println("Env Setup")
}
