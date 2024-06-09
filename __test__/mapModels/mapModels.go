package main

import (
	"encoding/json"
	"fmt"
	"github.com/crossevol/sqlc-model-codegen/codegen"
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
	groupedStructMetaMap := make(codegen.StructMetasMap)
	err = json.Unmarshal(bytes, &groupedStructMetaMap)
	if err != nil {
		log.Fatal(err)
	}

	dataMetas := codegen.Map2DataMetas(groupedStructMetaMap)

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
