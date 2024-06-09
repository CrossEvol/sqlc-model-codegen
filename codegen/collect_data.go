package codegen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// GetPackageName reads a *.go file from the given file path and returns the package name
func GetPackageName(filePath string) (string, error) {
	// Create a new token file set (required by the parser)
	fset := token.NewFileSet()

	// Parse the file
	node, err := parser.ParseFile(fset, filePath, nil, parser.PackageClauseOnly)
	if err != nil {
		return "", err
	}

	// Return the package name
	return node.Name.Name, nil
}

func CollectStructMetas(dir string) ([]*StructMeta, error) {
	// Get list of files
	files, err := getFilesInDirectory(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	// Print file names
	for _, file := range files {
		fmt.Println(file)
	}

	var structMetas []*StructMeta
	for _, file := range files {
		structMetas = append(structMetas, getStructTypesFromFile(file)...)
	}
	return structMetas, nil
}

func getFilesInDirectory(dir string) ([]string, error) {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func getStructTypesFromFile(filepath string) []*StructMeta {
	var structMetas []*StructMeta
	// Create a new token file set
	fset := token.NewFileSet()

	// Parse the .go file
	node, err := parser.ParseFile(fset, filepath, nil, parser.ParseComments)
	if err != nil {
		log.Fatalf("Failed to parse file: %v", err)
	}

	// Traverse the AST and find struct type declarations
	ast.Inspect(node, func(n ast.Node) bool {
		// Look for type declarations
		typeSpec, ok := n.(*ast.TypeSpec)
		if !ok {
			return true
		}

		// Look for struct types
		structType, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return true
		}

		structMeta := &StructMeta{}
		// Print the struct name
		fmt.Printf("Struct: %s\n", typeSpec.Name.Name)
		structMeta.Name = typeSpec.Name.Name

		// Iterate through the fields of the struct
		var fieldMetas []*FieldMeta
		for _, field := range structType.Fields.List {
			fieldMeta := &FieldMeta{}
			// Get the field names
			for _, name := range field.Names {
				// Print field name
				fmt.Printf("  Field: %s\n", name.Name)
				fieldMeta.Name = name.Name
			}

			// Print the field type
			fmt.Printf("    Type: %s\n", exprToString(field.Type))
			fieldMeta.Type = exprToString(field.Type)

			// Print the field tag (if any)
			if field.Tag != nil {
				fmt.Printf("    Tag: %s\n", field.Tag.Value)
				fieldMeta.Tag = field.Tag.Value
			}
			fieldMetas = append(fieldMetas, fieldMeta)

		}
		structMeta.FieldMetas = fieldMetas

		packageName, err := GetPackageName(filepath)
		if err != nil {
			log.Fatal(err)
		}
		structMeta.Package = packageName

		structMetas = append(structMetas, structMeta)
		return false
	})

	return structMetas
}

// Helper function to convert expression to string
func exprToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return exprToString(t.X) + "." + t.Sel.Name
	case *ast.StarExpr:
		return "*" + exprToString(t.X)
	case *ast.ArrayType:
		return "[]" + exprToString(t.Elt)
	case *ast.MapType:
		return "map[" + exprToString(t.Key) + "]" + exprToString(t.Value)
	case *ast.ChanType:
		var dir string
		switch t.Dir {
		case ast.SEND:
			dir = "chan<- "
		case ast.RECV:
			dir = "<-chan "
		default:
			dir = "chan "
		}
		return dir + exprToString(t.Value)
	default:
		return fmt.Sprintf("%T", expr)
	}
}

var BlackList = []string{"Queries"}

func GroupStructMetas(structMetas []*StructMeta) ([]*StructMeta, StructMetasMap, error) {
	for i, structMeta := range structMetas {
		for _, s := range BlackList {
			if s == structMeta.Name {
				structMetas = append(structMetas[:i], structMetas[i+1:]...)
				break
			}
		}
	}

	// filter out Name without Params$
	var plainStructMetas []*StructMeta
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
	groupedStructMetaMap := make(StructMetasMap)
	for _, structMeta := range plainStructMetas {
		groupedStructMetaMap[structMeta.Name] = []*StructMeta{structMeta}
	}
	for key, value := range groupedStructMetaMap {
		createParams := fmt.Sprintf("Create%sParams", key)
		updateParams := fmt.Sprintf("Update%sParams", key)
		for _, structMeta := range structMetas {
			if createParams == structMeta.Name || updateParams == structMeta.Name {
				value = append(value, structMeta)
				groupedStructMetaMap[key] = value
			}
		}
	}

	// copy the Type from [Model] to [UpdateModelParams], if not have [UpdateModelParams], should pass
	// only apply for sqlite generated code, because it will use interface{} but not *Type or sql.NullString
	for key, metas := range groupedStructMetaMap {
		var target *StructMeta
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

		var origin *StructMeta
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
	return plainStructMetas, groupedStructMetaMap, nil
}

func Map2DataMetas(groupedStructMetaMap StructMetasMap) []*DataMeta {
	var dataMetas []*DataMeta

	for key, structMetas := range groupedStructMetaMap {
		var dataMeta DataMeta
		for _, structMeta := range structMetas {
			if structMeta.Name == key {
				dataMeta.PlainModel = *structMeta
			} else if structMeta.Name == fmt.Sprintf("Create%sParams", key) {
				dataMeta.CreateModel = *structMeta
			} else if structMeta.Name == fmt.Sprintf("Update%sParams", key) {
				dataMeta.UpdateModel = *structMeta
			}
			if dataMeta.Package == "" {
				dataMeta.Package = structMeta.Package
			}
		}
		dataMetas = append(dataMetas, &dataMeta)
	}
	return dataMetas
}
