package main

import (
	"fmt"
	"github.com/crossevol/sqlc-model-codegen/codegen"
	"log"
)

func main() {
	// Example usage
	filePath := "D:\\GOLANG_CODE\\sqlc-model-codegen\\codegen\\codegen_test.go"
	pkgName, err := codegen.GetPackageName(filePath)
	if err != nil {
		log.Fatalf("Error getting package name: %v", err)
	}

	fmt.Printf("Package name: %s\n", pkgName)
}
