package io

// Model holds an entire parsed Archie model
type Model struct {
	Elements     []Element     `json:""`
	Associations []Association `json:""`
}

// Association holds a parsed Archie association
type Association struct {
	Source      string   `json:""`
	Destination string   `json:""`
	Tags        []string `json:"omitempty"`
}

// Element holds a parsed Archie element (actor or item)
type Element struct {
	Name         string        `json:""`
	Kind         string        `json:"omitempty"`
	Tags         []string      `json:"omitempty"`
	Children     []interface{} `json:"omitempty"`
	Associations []Association `json:"omitempty"`
}
