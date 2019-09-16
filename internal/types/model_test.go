package types

import (
	"testing"
)

// Test building a simple model
func TestModel(t *testing.T) {
	// Create a test model
	GetModel()
}

func GetModel() Model {
	// First create a model
	m := Model{}

	// Create some items
	system := NewItem("System")
	ux := NewItem("UX")
	fg := NewItem("Force generation")
	ft := NewItem("Force transmission")

	// Add some children to the system
	system.AddChild(ux)
	system.AddChild(fg)
	system.AddChild(ft)

	// Add some children to the sub-systems
	ux.AddChild(NewItem("Seat"))
	ux.AddChild(NewItem("Windshield"))
	ux.AddChild(NewItem("Windshield"))
	fuelTank := NewItem("Fuel tank")
	motor := NewItem("Motor")
	fg.AddChild(fuelTank)
	fg.AddChild(motor)
	wheel := NewItem("Wheel")
	axle := NewItem("Axle")
	ft.AddChild(wheel)
	ft.AddChild(axle)

	// Finally, add the system to the model
	m.AddElement(system)

	// Add some relationships
	m.AddRelationship(wheel, axle)
	m.AddRelationship(fuelTank, motor)

	// Return model
	return m
}
