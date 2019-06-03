package flex

// DataProvider interface
type DataProvider interface {
	GetOutputComponent() OutputComponent
	Models() (interface{}, error)
}

func getOutputComponent(p DataProvider) OutputComponent {
	return p.GetOutputComponent()
}

func NewDataProvider(cs ComponentSettings, typ string) (d DataProvider) {
	switch typ {
	case "book":
		d = NewBookDataProvider(cs)
	case "news":
		d = NewNewsDataProvider(cs)
	case "proposed_list":
		d = NewProposedListDataProvider(cs)
	case "Publisher":
		d = NewPublisherDataProvider(cs)
	}

	return
}
