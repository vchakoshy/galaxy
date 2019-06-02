package flex

// DataProvider interface
type DataProvider interface {
	getOutputComponent() OutputComponent
}

func getOutputComponent(p DataProvider) OutputComponent {
	return p.getOutputComponent()
}
