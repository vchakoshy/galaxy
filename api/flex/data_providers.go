package flex

// DataProvider interface
type DataProvider interface {
	getOutputComponent(ComponentSettings, string) OutputComponent
}

func getOutputComponent(p DataProvider, c ComponentSettings, t string) OutputComponent {
	return p.getOutputComponent(c, t)
}
