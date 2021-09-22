package models


type StaticProfile struct {
	Metadata struct {
		Date float32 `json:"date"`
		Type string  `json:"type"`
		Uid  string  `json:"uid"`
	} `json:"metadata"`
	Profile struct {
		Classes map[string]Class `json:"classes"`
	} `json:"profile"`
}

type Class struct {
	Methods map[string]Method `json:"methods"`
	Constructors []Constructor `json:"constructors"`
}

type Method struct {
	Overloads []Overload `json:"overloads"`
}

type Constructor struct {
	Arguments []TypeMap `json:"arguments"`
}

type Overload struct {
	Arguments []TypeMap `json:"arguments"`
	ReturnType TypeMap `json:"returnType"`
}

type TypeMap struct {
	ClassName string `json:"className"`
	Name string `json:"name"`
	Type string `json:"type"`
	Size uint32 `json:"size"`
	ByteSize uint32 `json:"byteSize"`
	DefaultValue interface{} `json:"defaultValue"`
}