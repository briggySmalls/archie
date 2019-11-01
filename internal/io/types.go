package io

type Model struct {
	Elements     []Element     `json:"elements"`
	Associations []Association `json:"associations"`
}

type Element struct {
	Name       string        `json:"name"`
	Kind       string        `json:"kind,omitempty"`
	Tags []string        `json:"technology,omitempty"`
	Children   []interface{} `json:"children,omitempty"`
}

type ElementWithChildren struct {
	Name       string        `json:"name"`
	Kind       string        `json:"kind,omitempty"`
	Tags []string        `json:"technology,omitempty,flow"`
	Children   []interface{} `json:"children,omitempty"`
}

type Association struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
}
