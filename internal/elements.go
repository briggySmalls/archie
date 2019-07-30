type Item struct {
	Name     string
	Children []*Item
}

type actor struct {
	Name string
}

// Indicate items are not actors
func (e *Item) IsActor() bool {
	return false
}

// Indicate actors are actors
func (e *Actor) IsActor() bool {
	return true
}
