package codegen

import (
	"fmt"
	"github.com/crossevol/sqlc-model-codegen/internal/util"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

const DtoTemplate = `
package {{ .DestPackage }}

{{if IsNotEmpty .Imports}}import (
	{{- range $_,$import := .Imports }}
		"{{ $import }}"
	{{- end }}
){{ end }}

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
	{{range $index,$fieldMeta := .CreateModel.FieldMetas }}
		{{- if OfType $fieldMeta.Type}}
			{{- $fieldMeta.Name }}:= {{ Convert2InExpr $fieldMeta.Type $fieldMeta.Name }}
		{{ end -}}
	{{end}}
	params := {{ .Package }}.{{ .CreateModel.Name | ToCamel }}Params{
        {{range $index,$fieldMeta := .CreateModel.FieldMetas }}
            {{- $fieldMeta.Name }}: {{if OfType $fieldMeta.Type}} {{if IsPtr $fieldMeta.Type}}&{{end}}{{ $fieldMeta.Name }} {{else}} dto.{{ $fieldMeta.Name }} {{ end }} ,
        {{end}}
    }
    return &params
}
{{end}}

{{if IsNotEmpty .UpdateModel.Name }}
func (dto *{{ .UpdateModel.Name | ToCamel }}DTO)Map2{{ .UpdateModel.Name | ToCamel }}Params() *{{ .Package }}.{{ .UpdateModel.Name | ToCamel }}Params {
	{{range $index,$fieldMeta := .CreateModel.FieldMetas }}
		{{- if OfType $fieldMeta.Type}}
			{{- $fieldMeta.Name }}:= {{ Convert2InExpr $fieldMeta.Type $fieldMeta.Name }}
		{{ end -}}
	{{end}}
	params := {{ .Package }}.{{ .UpdateModel.Name | ToCamel }}Params{
        {{range $index,$fieldMeta := .UpdateModel.FieldMetas }}
            {{- $fieldMeta.Name }}: {{if OfType $fieldMeta.Type}} {{if IsPtr $fieldMeta.Type}}&{{end}}{{ $fieldMeta.Name }} {{else}} dto.{{ $fieldMeta.Name }} {{ end }} ,
        {{end}}
    }
    return &params
}
{{end}}

{{if IsNotEmpty .PlainModel.Name }}
func Map2{{ .PlainModel.Name | ToCamel }}DTO(entity *{{ .Package }}.{{ .PlainModel.Name | ToCamel }}) *{{ .PlainModel.Name | ToCamel }}DTO {
	dto := {{ .PlainModel.Name | ToCamel }}DTO{
        {{range $index,$fieldMeta := .PlainModel.FieldMetas }}
            {{- $fieldMeta.Name }}: {{if OfType $fieldMeta.Type}} {{ Convert2OutExpr $fieldMeta.Type $fieldMeta.Name }} {{else}} entity.{{ $fieldMeta.Name }} {{ end }} ,
        {{end}}
    }
    return &dto
}
{{end}}
`

func CrudGen(dataMetas DataMetas, destDir string, imports []string) error {
	lastIndex := strings.LastIndex(destDir, string(filepath.Separator))
	destPackage := destDir[lastIndex+1:]

	tmpl := template.Must(template.New("CrudTemplate").Funcs(util.TemplateFuncMap()).Parse(DtoTemplate))

	for _, dataMeta := range dataMetas {
		dataMeta.DestPackage = destPackage
		dataMeta.Imports = imports
		if kDebug {
			fmt.Println(dataMeta.PlainModel.Name)
		}
		var name = dataMeta.PlainModel.Name
		var content strings.Builder
		err := tmpl.Execute(&content, dataMeta)
		if err != nil {
			return err
		}
		if _, err := os.Stat(destDir); err != nil {
			os.Mkdir(destDir, fs.ModePerm)
		}
		if err := os.WriteFile(filepath.Join(destDir, fmt.Sprintf("%s.dto.go", name)), []byte(content.String()), os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
