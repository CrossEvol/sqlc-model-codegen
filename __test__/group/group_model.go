package main

import (
	"encoding/json"
	"fmt"
	"github.com/crossevol/sqlc-model-codegen/codegen"
	"log"
	"os"
)

const PlainStructsFile = "plain_structs.json"
const GroupedStructsFile = "grouped_structs.json"

func main() {
	bytes, err := os.ReadFile("struct_metas.json")
	if err != nil {
		log.Fatal(err)
	}
	var structMetas []*codegen.StructMeta
	err = json.Unmarshal(bytes, &structMetas)
	if err != nil {
		log.Fatal(err)
	}

	plainStructMetas, groupedStructMetaMap, err := codegen.GroupStructMetas(structMetas)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Create or open the PlainStructsFile
	file, err := os.Create(PlainStructsFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Encode data to JSON format
	encoder := json.NewEncoder(file)
	err = encoder.Encode(plainStructMetas)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println(fmt.Sprintf("JSON data written to %s successfully", PlainStructsFile))

	// Create or open the GroupedStructsFile
	file, err = os.Create(GroupedStructsFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Encode data to JSON format
	encoder = json.NewEncoder(file)
	err = encoder.Encode(groupedStructMetaMap)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println(fmt.Sprintf("JSON data written to %s successfully", GroupedStructsFile))
}
