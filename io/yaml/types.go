package yaml

type Model struct {
	Elements     []Element
	Associations []Association
}

type Element struct {
	Name       string
	Type       string        `yaml:",omitempty"`
	Technology string        `yaml:",omitempty"`
	Children   []interface{} `yaml:",omitempty"`
}

type Association struct {
	Source      string
	Destination string
}
