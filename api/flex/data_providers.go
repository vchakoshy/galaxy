package flex

// DataProvider interface
type DataProvider interface {
	getOutputComponent() OutputComponent
	// Models() (interface{}, error)
}

func getOutputComponent(p DataProvider) OutputComponent {
	return p.getOutputComponent()
}
