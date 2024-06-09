package main

import (
	"encoding/json"
	"fmt"
	"github.com/crossevol/sqlc-model-codegen/__test__/gm"
	"github.com/crossevol/sqlc-model-codegen/internal/util"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const DtoTemplate = `
package curdGen

{{if IsNotEmpty .PlainModel.Name }}
type {{ .PlainModel.Name | ToCamel }}DTO struct {
  {{range $index,$fieldMeta := .PlainModel.FieldMetas }}
      {{- $fieldMeta.Name -}}
      {{if OfType $fieldMeta.Type}} {{ ConvertType $fieldMeta.Type }} {{else}} {{ $fieldMeta.Type }} {{ end -}}
      {{ $fieldMeta.Tag }}
  {{ end -}}
}
{{end}}

{{if IsNotEmpty .CreateModel.Name }}
type {{ .CreateModel.Name | ToCamel }}DTO struct {
  {{range $index,$fieldMeta := .CreateModel.FieldMetas }}
      {{- $fieldMeta.Name -}}
      {{if OfType $fieldMeta.Type}} {{ ConvertType $fieldMeta.Type }} {{else}} {{ $fieldMeta.Type }} {{ end -}}
      {{ $fieldMeta.Tag }}
  {{ end -}}
}
{{end}}

{{if IsNotEmpty .UpdateModel.Name }}
type {{ .UpdateModel.Name | ToCamel }}DTO struct {
  {{range $index,$fieldMeta := .UpdateModel.FieldMetas }}
      {{- $fieldMeta.Name -}}
      {{if OfType $fieldMeta.Type}} {{ ConvertType $fieldMeta.Type }} {{else}} {{ $fieldMeta.Type }} {{ end -}}
      {{ $fieldMeta.Tag }}
  {{ end -}}
}
{{end}}


{{if IsNotEmpty .CreateModel.Name }}
func (dto *{{ .CreateModel.Name | ToCamel }}DTO)Map2{{ .CreateModel.Name | ToCamel }}Params() *{{ .Package }}.{{ .CreateModel.Name | ToCamel }}Params {
    return &{{ .Package }}.{{ .CreateModel.Name | ToCamel }}Params{
        {{range $index,$fieldMeta := .CreateModel.FieldMetas }}
            {{- $fieldMeta.Name }}: {{if OfType $fieldMeta.Type}} {{ Convert2InExpr $fieldMeta.Type $fieldMeta.Name }} {{else}} dto.{{ $fieldMeta.Name }} {{ end }} ,
        {{end}}
    }
}
{{end}}

{{if IsNotEmpty .UpdateModel.Name }}
func (dto *{{ .UpdateModel.Name | ToCamel }}DTO)Map2{{ .UpdateModel.Name | ToCamel }}Params() *{{ .Package }}.{{ .UpdateModel.Name | ToCamel }}Params {
    return &{{ .Package }}.{{ .UpdateModel.Name | ToCamel }}Params{
        {{range $index,$fieldMeta := .UpdateModel.FieldMetas }}
            {{- $fieldMeta.Name }}: {{if OfType $fieldMeta.Type}} {{ Convert2InExpr $fieldMeta.Type $fieldMeta.Name }} {{else}} dto.{{ $fieldMeta.Name }} {{ end }} ,
        {{end}}
    }
}
{{end}}

{{if IsNotEmpty .PlainModel.Name }}
func (entity *{{ .Package }}.{{ .PlainModel.Name | ToCamel }})Map2{{ .PlainModel.Name | ToCamel }}DTO() *{{ .PlainModel.Name | ToCamel }}DTO {
    return &{{ .PlainModel.Name | ToCamel }}DTO{
        {{range $index,$fieldMeta := .PlainModel.FieldMetas }}
            {{- $fieldMeta.Name }}: {{if OfType $fieldMeta.Type}} {{ Convert2OutExpr $fieldMeta.Type $fieldMeta.Name }} {{else}} entity.{{ $fieldMeta.Name }} {{ end }} ,
        {{end}}
    }
}
{{end}}
`

func main() {
	gen := "curdGen"
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	bytes, err := os.ReadFile("data_metas.json")
	if err != nil {
		log.Fatal(err)
	}

	var dataMetas gm.DataMetas
	err = json.Unmarshal(bytes, &dataMetas)
	if err != nil {
		log.Fatal(err)
	}

	tmpl := template.Must(template.New("CrudTemplate").Funcs(util.TemplateFuncMap()).Parse(DtoTemplate))

	for _, dataMeta := range dataMetas {
		fmt.Println(dataMeta.PlainModel.Name)
		var name = dataMeta.PlainModel.Name
		var content strings.Builder
		err = tmpl.Execute(&content, dataMeta)
		if err != nil {
			log.Fatal(err)
		}
		if _, err := os.Stat(filepath.Join(wd, gen)); err != nil {
			os.Mkdir(filepath.Join(wd, gen), fs.ModePerm)
		}
		if err := os.WriteFile(filepath.Join(wd, gen, fmt.Sprintf("%s.dto.go", name)), []byte(content.String()), os.ModePerm); err != nil {
			fmt.Println(err)
		}
		//fmt.Println(content.String())
	}

}
