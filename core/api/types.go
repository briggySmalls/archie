package api

type Model struct {
	Elements     []Element
	Associations []Association
}

type Element struct {
	Name       string
	Kind       string        `yaml:",omitempty" json:",omitempty"`
	Technology string        `yaml:",omitempty" json:",omitempty"`
	Children   []interface{} `yaml:",omitempty" json:",omitempty"`
}

type ElementWithChildren struct {
	Name       string
	Kind       string        `yaml:",omitempty" json:",omitempty"`
	Technology string        `yaml:",omitempty" json:",omitempty"`
	Children   []interface{} `yaml:",omitempty" json:",omitempty"`
}

type Association struct {
	Source      string
	Destination string
}
