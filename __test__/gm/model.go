package gm

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
