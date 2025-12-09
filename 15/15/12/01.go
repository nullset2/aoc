package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var G = 0.0

func traverseJSON(data interface{}, indent string) {
	switch v := data.(type) {
	case map[string]interface{}:
		for _, value := range v {
			traverseJSON(value, indent+"    ")
		}
	case []interface{}:
		for _, value := range v {
			traverseJSON(value, indent+"    ")
		}
	case float64:
		fmt.Printf("%sNumber: %f\n", indent, v)
		G += v
	}
}

func main() {
	jsonFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	var dat map[string]interface{}
	bytes, _ := ioutil.ReadAll(jsonFile)
	if err := json.Unmarshal(bytes, &dat); err != nil {
		panic(err)
	}
	traverseJSON(dat, "")
	fmt.Println(G)
}
