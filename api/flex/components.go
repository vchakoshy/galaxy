package flex

func handleListComponent(cs ComponentSettings, t string) OutputComponent {

	switch cs.Settings.DataProvider {
	case "BOOK":
		var com BookDataProvider
		return com.getOutputComponent(cs, t)

	case "PROPOSED_LIST":
		var com ProposedListDataProvider
		return com.getOutputComponent(cs, t)

	case "PUBLISHER":
		var com PublisherDataProvider
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
