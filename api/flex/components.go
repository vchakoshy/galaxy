package flex

const (
	dataProviderTypeBook         = "BOOK"
	dataProviderTypeProposedList = "PROPOSED_LIST"
	dataproviderTypePublisher    = "PUBLISHER"
	dataproviderTypeNews         = "NEWS"
)

func handleListComponent(cs ComponentSettings, t string) OutputComponent {
	switch cs.Settings.DataProvider {
	case dataProviderTypeBook:
		return getOutputComponent(BookDataProvider{}, cs, t)
	case dataProviderTypeProposedList:
		return getOutputComponent(ProposedListDataProvider{}, cs, t)
	case dataproviderTypePublisher:
		return getOutputComponent(PublisherDataProvider{}, cs, t)
	case dataproviderTypeNews:
		return getOutputComponent(NewsDataProvider{}, cs, t)
	}

	return OutputComponent{}
}

const (
	resourceTypeCustom = "CUSTOM"
)

const (
	componentTypeHlBooksLibrary = "HL_BOOKS_LIBRARY"
)

func handleSingleComponent(cs ComponentSettings, t string) (com OutputComponent) {
	com.Title = cs.Elements.Title.Value.Static
	com.Icon = cs.Elements.Icon.Value
	com.Type = t
	com.ResourceType = resourceTypeCustom

	g := Generic{
		Image: cs.Settings.ChildElements.Image.Value.Static,
		Ratio: cs.Settings.ChildElements.Ratio.Value,
	}
	a := getAction(cs.Settings.ChildElements.Action.Action)
	if a.Type != "" {
		g.Action = a
	}
	com.Data.Items.Generic = append(com.Data.Items.Generic, g)
	return
}

func handleBooksLibraryComponent(cs ComponentSettings) (com OutputComponent) {
	var g Generic
	g.Format = cs.Settings.Format
	g.ContentType = cs.Settings.ContentType
	com.Data.Items.Generic = append(com.Data.Items.Generic, g)
	com.Title = cs.Elements.Title.Value.Static
	com.Icon = cs.Elements.Icon.Value
	com.Type = componentTypeHlBooksLibrary
	com.ResourceType = resourceTypeCustom
	return
}
