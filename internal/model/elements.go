package model

import (
	"fmt"
)

const (
	actor = iota
	item
)

type element struct {
	name string
	tags []string
	kind uint
}

// Element describes a model element
type Element interface {
	Name() string
	Tags() []string
	ID() string
	IsActor() bool
	HasTag(tag string) bool
}

// NewItem creates a new item element
func NewItem(name string, tags []string) Element {
	el := newElement(item)
	el.name = name
	el.tags = tags
	return &el
}

// NewActor creates a new actor element
func NewActor(name string) Element {
	el := newElement(actor)
	el.name = name
	return &el
}

func newElement(kind uint) element {
	return element{
		kind: kind,
	}
}

func (e *element) ID() string {
	return fmt.Sprintf("%p", e)
}

func (e *element) IsActor() bool {
	return e.kind == actor
}

func (e *element) HasTag(tag string) bool {
	for _, t := range e.Tags() {
		if t == tag {
			return true
		}
	}
	return false
}

func (e *element) Name() string {
	return e.name
}

func (e *element) Tags() []string {
	return e.tags
}
