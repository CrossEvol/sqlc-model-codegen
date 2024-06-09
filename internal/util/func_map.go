package util

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"github.com/jinzhu/inflection"
	"regexp"
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

func Convert2InExpr(fieldType string, fieldName string) string {
	expr := fieldName
	switch fieldType {
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
	default:
		expr = fmt.Sprintf("dto.%s", expr)
	}
	return expr
}

func Convert2OutExpr(fieldType string, fieldName string) string {
	expr := fieldName
	switch fieldType {
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
	default:
		expr = fmt.Sprintf("entity.%s", expr)
	}
	return expr
}

func IsEmpty(v string) bool {
	return len(v) == 0
}

func IsNotEmpty(v string) bool {
	return len(v) != 0
}

func IsPtr(v string) bool {
	// Regular expression pattern
	pattern := regexp.MustCompile(`^\*`)
	return pattern.MatchString(v)
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
		"OfType":           OfType,
		"ConvertType":      ConvertType,
		"Convert2InExpr":   Convert2InExpr,
		"Convert2OutExpr":  Convert2OutExpr,
		"IsEmpty":          IsEmpty,
		"IsNotEmpty":       IsNotEmpty,
		"IsPtr":            IsPtr,
	}
	return funcMap
}
