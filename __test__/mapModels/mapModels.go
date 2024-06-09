package main

import (
	"encoding/json"
	"fmt"
	"github.com/crossevol/sqlc-model-codegen/__test__/gm"
	"log"
	"os"
)

const DataMetasFile = "data_metas.json"

func main() {
	bytes, err := os.ReadFile("grouped_structs.json")
	if err != nil {
		log.Fatal(err)
	}
	// group [Post, CreatePostParams, UpdatePostParams]
	groupedStructMetaMap := make(gm.StructMetasMap)
	err = json.Unmarshal(bytes, &groupedStructMetaMap)
	if err != nil {
		log.Fatal(err)
	}

	var dataMetas []*gm.DataMeta

	for key, structMetas := range groupedStructMetaMap {
		var dataMeta gm.DataMeta
		for _, structMeta := range structMetas {
			if structMeta.Name == key {
				dataMeta.PlainModel = structMeta
			} else if structMeta.Name == fmt.Sprintf("Create%sParams", key) {
				dataMeta.CreateModel = structMeta
			} else if structMeta.Name == fmt.Sprintf("Update%sParams", key) {
				dataMeta.UpdateModel = structMeta
			}
		}
		dataMetas = append(dataMetas, &dataMeta)
	}

	// Create or open the PlainStructsFile
	file, err := os.Create(DataMetasFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Encode data to JSON format
	encoder := json.NewEncoder(file)
	err = encoder.Encode(dataMetas)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	fmt.Println(fmt.Sprintf("JSON data written to %s successfully", DataMetasFile))
}
