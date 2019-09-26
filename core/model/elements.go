package model

import (
	"fmt"
)

const (
	actor = iota
	item
)

type Element struct {
	Name string
	Kind uint
}

// Create a new item
func NewItem(name string) Element {
	el := newElement(item)
	el.Name = name
	return el
}

// Create a new actor
func NewActor(name string) Element {
	el := newElement(actor)
	el.Name = name
	return el
}

// Create an element
func newElement(kind uint) Element {
	return Element{
		Kind: kind,
	}
}

func (e *Element) Id() string {
	return fmt.Sprintf("%p", e)
}
