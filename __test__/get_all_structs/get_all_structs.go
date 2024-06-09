package main

import (
	"encoding/json"
	"fmt"
	"github.com/crossevol/sqlc-model-codegen/__test__/gm"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir := "internal/database/sqliteDao"

	// Get list of files
	files, err := getFilesInDirectory(dir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print file names
	for _, file := range files {
		fmt.Println(file)
	}

	var structMetas []*gm.StructMeta
	for _, file := range files {
		structMetas = append(structMetas, getStructTypesFromFile(file)...)
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

func getStructTypesFromFile(filepath string) []*gm.StructMeta {
	var structMetas []*gm.StructMeta

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

		structMeta := &gm.StructMeta{}
		// Print the struct name
		fmt.Printf("Struct: %s\n", typeSpec.Name.Name)
		structMeta.Name = typeSpec.Name.Name

		// Iterate through the fields of the struct
		var fieldMetas []*gm.FieldMeta
		for _, field := range structType.Fields.List {
			fieldMeta := &gm.FieldMeta{}
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
