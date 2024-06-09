package main

import (
	"encoding/json"
	"fmt"
	"github.com/crossevol/sqlc-model-codegen/__test__/gm"
	"log"
	"os"
	"regexp"
	"strings"
)

const PlainStructsFile = "plain_structs.json"
const GroupedStructsFile = "grouped_structs.json"

var BlackList = []string{"Queries"}

func main() {
	bytes, err := os.ReadFile("struct_metas.json")
	if err != nil {
		log.Fatal(err)
	}
	var structMetas []*gm.StructMeta
	err = json.Unmarshal(bytes, &structMetas)
	if err != nil {
		log.Fatal(err)
	}

	for i, structMeta := range structMetas {
		for _, s := range BlackList {
			if s == structMeta.Name {
				structMetas = append(structMetas[:i], structMetas[i+1:]...)
				break
			}
		}
	}

	// filter out Name without Params$
	var plainStructMetas []*gm.StructMeta
	for _, structMeta := range structMetas {
		reg, err := regexp.Compile("Params$")
		if err != nil {
			log.Fatal(err)
		}
		if !reg.MatchString(structMeta.Name) {
			plainStructMetas = append(plainStructMetas, structMeta)
		}
	}

	// group [Post, CreatePostParams, UpdatePostParams]
	groupedStructMetaMap := make(map[string][]*gm.StructMeta)
	for _, structMeta := range plainStructMetas {
		groupedStructMetaMap[structMeta.Name] = []*gm.StructMeta{structMeta}
	}
	for key, value := range groupedStructMetaMap {
		createParams := fmt.Sprintf("Create%sParams", key)
		updateParams := fmt.Sprintf("Update%sParams", key)
		for _, structMeta := range structMetas {
			if err != nil {
				log.Fatal(err)
			}
			if createParams == structMeta.Name || updateParams == structMeta.Name {
				value = append(value, structMeta)
				groupedStructMetaMap[key] = value
			}
		}
	}

	// copy the Type from [Model] to [UpdateModelParams], if not have [UpdateModelParams], should pass
	// only apply for sqlite generated code, because it will use interface{} but not *Type or sql.NullString
	for key, metas := range groupedStructMetaMap {
		var target *gm.StructMeta
		for _, meta := range metas {
			compile, err := regexp.Compile("^Update.*Params$")
			if err != nil {
				log.Fatal(err)
			}
			if compile.MatchString(meta.Name) {
				target = meta
				break
			}
		}

		var origin *gm.StructMeta
		for _, meta := range metas {
			if key == meta.Name {
				origin = meta
				break
			}
		}
		if target != nil {
			fmt.Println(target.Name)
		}
		if origin != nil {
			fmt.Println(origin.Name)
		}

		if target != nil && origin != nil {
			for _, targetFieldMeta := range target.FieldMetas {
				for _, originFieldMeta := range origin.FieldMetas {
					if targetFieldMeta.Name == originFieldMeta.Name && targetFieldMeta.Type != originFieldMeta.Type && targetFieldMeta.Type == "*ast.InterfaceType" {
						targetFieldMeta.Type = originFieldMeta.Type
						if strings.Index(targetFieldMeta.Type, "*") == -1 {
							targetFieldMeta.Type = "*" + targetFieldMeta.Type
						}
					}
				}
			}
		}

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
