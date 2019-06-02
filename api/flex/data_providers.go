package flex

// DataProvider interface
type DataProvider interface {
	GetOutputComponent() OutputComponent
	// Models() (interface{}, error)
}

func getOutputComponent(p DataProvider) OutputComponent {
	return p.GetOutputComponent()
}

func NewDataProviderByComponentSettings(cs ComponentSettings, typ string) DataProvider {
	var d DataProvider

	switch typ {
	case "book":
		d = NewBookDataProvider(cs)
	}

	return d

}
