package codegen

type StructMeta struct {
	Name       string       `json:"name"`
	FieldMetas []*FieldMeta `json:"field_meta"`
	Package    string       `json:"package"`
}

type FieldMeta struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Tag  string `json:"tag"`
}

type DataMeta struct {
	Package     string     `json:"package"`
	DestPackage string     `json:"dest_package"`
	Imports     []string   `json:"imports"`
	PlainModel  StructMeta `json:"plain_model"`
	CreateModel StructMeta `json:"create_model"`
	UpdateModel StructMeta `json:"update_model"`
}

type DataMetas = []*DataMeta

type StructMetasMap = map[string][]*StructMeta
