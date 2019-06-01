package flex

func handleListComponent(cs ComponentSettings, t string) OutputComponent {

	switch cs.Settings.DataProvider {
	case "BOOK":
		return getOutputComponent(BookDataProvider{}, cs, t)
	case "PROPOSED_LIST":
		return getOutputComponent(ProposedListDataProvider{}, cs, t)
	case "PUBLISHER":
		return getOutputComponent(PublisherDataProvider{}, cs, t)
	case "NEWS":
		var com NewsDataProvider
		return com.getOutputComponent(cs, t)
	}

	return OutputComponent{}
}

func handleSingleComponent(cs ComponentSettings, t string) (com OutputComponent) {
	com.Title = cs.Elements.Title.Value.Static
	com.Icon = cs.Elements.Icon.Value
	com.Type = t
	com.ResourceType = "CUSTOM"

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
	com.Type = "HL_BOOKS_LIBRARY"
	com.ResourceType = "CUSTOM"
	return
}
