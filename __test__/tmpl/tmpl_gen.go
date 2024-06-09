package main

import (
	"encoding/json"
	"fmt"
	"github.com/crossevol/sqlc-model-codegen/codegen"
	"log"
	"os"
	"path/filepath"
)

func main() {
	bytes, err := os.ReadFile("data_metas.json")
	if err != nil {
		log.Fatal(err)
	}

	var dataMetas codegen.DataMetas
	err = json.Unmarshal(bytes, &dataMetas)
	if err != nil {
		log.Fatal(err)
	}

	gen := "curdGen"
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	destDir := filepath.Join(wd, gen)

	codegen.CrudGen(dataMetas, destDir)

}
