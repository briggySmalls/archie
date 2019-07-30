func getModel() Model {
	// First create a model
	m := Model()

	// Add a few items
	lowLevelItems := []Item{
		{Name: "Seat"},
		{Name: "Windshield"},
		{Name: "Fuel Tank"},
		{Name: "Motor"},
		{Name: "Wheel"},
		{Name: "Axle"},
	}
	midLevelItems := []Item{
		{
			Name:     "UX",
			Children: {&lowLevelItems[0], &lowLevelItems[1]},
		},
		{
			Name:     "Force generation",
			Children: {&lowLevelItems[2], &lowLevelItems[3]},
		},
		{
			Name:     "Force transmission",
			Children: {&lowLevelItems[4], &lowLevelItems[5]},
		},
	}
	system := Item{
		Name:     "System",
		Children: {&midLevelItems[0], &midLevelItems[1], &midLevelItems[2]},
	}

	m.Items = append(m.Items, lowLevelItems...)
	m.Items = append(m.Items, midLevelItems...)
	m.Items = append(m.Items, system)

	// Add a driver
	driver := Actor{Name: "Driver"}
	m.Actors = append(m.Actors, &driver)

	// Add some relationships
	m.Relationships = append(m.Relationships, []Relationship{
		{ // Link wheel and axel
			Source: &lowLevelItems[2]
			Destination: &lowLevelItems[3],

		},
		{ // Link fuel tank and motor
			Source: &lowLevelItems[5]
			Destination: &lowLevelItems[5],

		},
	}...)

	// Return model
	return model
}
