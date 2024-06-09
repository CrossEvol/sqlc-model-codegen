package util

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"text/template"
)

func Quote(string string) string {
	return fmt.Sprintf("`%s`", string)
}

func Quote2(string string) string {
	return fmt.Sprintf("'%s'", string)
}

func LastFunc(index, length int) bool {
	return index == length-1
}

var SpecificTypeSet = []string{
	"*time.Time",
	"time.Time",
}

func OfType(typeName string) bool {
	var flag bool
	for _, s := range SpecificTypeSet {
		if s == typeName {
			flag = true
			break
		}
	}
	return flag
}

func ConvertType(typeName string) string {
	newType := typeName
	switch typeName {
	case "*time.Time":
		{
			newType = "int64"
			break
		}
	case "time.Time":
		{
			newType = "int64"
			break
		}
	}
	return newType
}

func Convert2InExpr(fieldName string) string {
	expr := fieldName
	switch fieldName {
	case "*time.Time":
		{
			expr = fmt.Sprintf("time.UnixMilli(dto.%s)", fieldName)
			break
		}
	case "time.Time":
		{
			expr = fmt.Sprintf("time.UnixMilli(dto.%s)", fieldName)
			break
		}
	}
	return expr
}

func Convert2OutExpr(fieldName string) string {
	expr := fieldName
	switch fieldName {
	case "*time.Time":
		{
			expr = fmt.Sprintf("entity.%s.UnixMilli()", fieldName)
			break
		}
	case "time.Time":
		{
			expr = fmt.Sprintf("entity.%s.UnixMilli()", fieldName)
			break
		}
	}
	return expr
}

var CreateList = []string{"Create", "create", "CREATE"}

func TemplateFuncMap() template.FuncMap {
	funcMap := template.FuncMap{
		"ToSnake":          strcase.ToSnake,
		"ToCamel":          strcase.ToCamel,
		"ToLower":          strcase.ToLowerCamel,
		"ToScreamingSnake": strcase.ToScreamingSnake,
		"Plural":           inflection.Plural,
		"Singular":         inflection.Singular,
		"Quote":            Quote,
		"Add":              func(a, b int) int { return a + b },
		"Last":             LastFunc,
	}
	return funcMap
}
