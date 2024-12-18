package services

import (
	"encoding/json"
	"fmt"
	"os"
)


func getFilePath(whichPath string) string {
	wd, err := os.Getwd()
	var filePath string
	if err != nil {
		fmt.Println("Error fetching the file", err)
	}
	if whichPath == "channels" {
		filePath = wd + "/data/channels.json"
	} else if whichPath == "messages" {
		filePath = wd + "/data/messages.json"
	} else {
		filePath = "NA"
	}
	return filePath
}

func readJSONFile(p_interface interface{}, d_type string) {
	/*
		p_interface: used to reference type interface json to store
		d_type: used to file which json file path you want to store to.
	*/
	filePath := getFilePath(d_type)
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(p_interface)

	if err != nil {
		fmt.Println("Error reading JSON from file:", err)
	}
}

func writeJSONFile(p_interface interface{}, d_type string) {
	filePath := getFilePath(d_type)
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	if err := encoder.Encode(p_interface); err != nil {
		fmt.Println("error encoding JSON: ", err)
	}

}


