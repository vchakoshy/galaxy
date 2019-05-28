package flex

func handleListComponent(cs ComponentSettings, t string) OutputComponent {

	switch cs.Settings.DataProvider {
	case "BOOK":
		com := DataProvidersBook{}
		return com.getGeneric(cs, t)

	case "PROPOSED_LIST":
		com := DataProvidersProposedList{}
		return com.getGeneric(cs, t)
	}

	return OutputComponent{}
}

func handleSingleComponent(cs ComponentSettings, t string) (com OutputComponent) {
	com.Title = cs.Elements.Title.Value.Static
	com.Icon = cs.Elements.Icon.Value
	com.Type = t
	com.ResourceType = "CUSTOM"

	m := make(map[string]interface{})

	m["image"] = cs.Settings.ChildElements.Image.Value.Static
	m["ratio"] = cs.Settings.ChildElements.Ratio.Value
	a := getAction(cs.Settings.ChildElements.Action.Action)
	if a.Type != "" {
		m["action"] = a
	}
	com.Data.Items.Generic = append(com.Data.Items.Generic, m)
	return
}

func handleBooksLibraryComponent(cs ComponentSettings) (com OutputComponent) {
	m := make(map[string]string)
	m["format"] = cs.Settings.Format
	m["content_type"] = cs.Settings.ContentType
	com.Data.Items.Generic = append(com.Data.Items.Generic, m)
	com.Title = cs.Elements.Title.Value.Static
	com.Icon = cs.Elements.Icon.Value
	com.Type = "HL_BOOKS_LIBRARY"
	com.ResourceType = "CUSTOM"
	return
}
