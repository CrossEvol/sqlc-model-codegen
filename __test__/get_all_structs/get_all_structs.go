package main

import (
	"encoding/json"
	"fmt"
	"github.com/crossevol/sqlc-model-codegen/codegen"
	"os"
)

func main() {
	dir := "internal/database/sqliteDao"

	structMetas, err := codegen.CollectStructMetas(dir)
	if err != nil {
		return
	}

	// Create or open the file
	file, err := os.Create("struct_metas.json")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Encode data to JSON format
	encoder := json.NewEncoder(file)
	err = encoder.Encode(structMetas)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println("JSON data written to file successfully")

}
