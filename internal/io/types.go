package io

// Model holds an entire parsed Archie model
type Model struct {
	Elements     []Element     `json:"elements"`
	Associations []Association `json:"associations"`
}

// Element holds a parsed Archie element (actor or item)
type Element struct {
	Name     string        `json:"name"`
	Kind     string        `json:"kind,omitempty"`
	Tags     []string      `json:"technology,omitempty"`
	Children []interface{} `json:"children,omitempty"`
}

// Association holds a parsed Archie association
type Association struct {
	Source      string   `json:"source"`
	Destination string   `json:"destination"`
	Tags        []string `json:"tag"`
}
